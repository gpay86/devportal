import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;

public class EmvicoSample {
    // Private constructor to prevent instantiation


    // Cấu trúc TLV đơn giản
    public static class EmvField {
        public String id;
        public String value;
        public Map<String, EmvField> children; // nếu đây là template lồng (38, 62, 64, hoặc field con kiểu template)

        public EmvField(String id, String value) {
            this.id = id;
            this.value = value;
            this.children = null;
        }

        @Override
        public String toString() {
            return "EmvField{id='" + id + "', value='" + value + "', children=" + children + "}";
        }
    }

    /**
     * Parse chuỗi EMV TLV từ đầu đến hết
     */
    public static List<EmvField> parseTlvs(String data) {
        List<EmvField> result = new ArrayList<>();
        int idx = 0;
        while (idx + 4 <= data.length()) {
            String id = data.substring(idx, idx + 2);
            int len = Integer.parseInt(data.substring(idx + 2, idx + 4));
            idx += 4;

            if (idx + len > data.length()) {
                throw new IllegalArgumentException("Invalid TLV length for id " + id);
            }

            String value = data.substring(idx, idx + len);
            idx += len;

            EmvField field = new EmvField(id, value);

            // Các field là template có cấu trúc TLV lồng bên trong
            if (id.equals("38") || id.equals("62") || id.equals("64")) {
                field.children = toMap(parseTlvs(value));
            }

            result.add(field);

            // Theo chuẩn EMV, CRC (63) là field cuối cùng
            if (id.equals("63")) {
                break;
            }
        }
        return result;
    }

    private static Map<String, EmvField> toMap(List<EmvField> list) {
        return list.stream().collect(Collectors.toMap(f -> f.id, Function.identity(), (a, b) -> a));
    }

    /**
     * Parse TLV con (sub template) generic
     */
    public static Map<String, EmvField> parseNestedTemplate(String value) {
        return toMap(parseTlvs(value));
    }

    /**
     * Decode VietQR string -> ObjectNode JSON
     */
    public static GpayFtToBankRequest decodeVietQrToJson(String qrString) {

        GpayFtToBankRequest gpayFtToBankRequest = new GpayFtToBankRequest();
        List<EmvField> rootFields = parseTlvs(qrString);
        Map<String, EmvField> rootMap = toMap(rootFields);


        // 38 - Merchant Account Information (VietQR)
        EmvField f38 = rootMap.get("38");
        if (f38 != null) {


            if (f38.children == null) {
                f38.children = parseNestedTemplate(f38.value);
            }
            Map<String, EmvField> maInfo = f38.children;

            // 00 - GUID (A000000727)


            // 01 - tổ chức thụ hưởng: bên trong là TLV chứa BIN + account
            EmvField orgField = maInfo.get("01");
            Map<String, EmvField> orgChildren = null;
            if (orgField != null) {
                try {
                    orgChildren = parseNestedTemplate(orgField.value);
                } catch (Exception ex) {
                    orgChildren = null;
                }

                if (orgChildren != null && !orgChildren.isEmpty()) {
                    // Sub 00 - Bank BIN
                    EmvField bankBinField = orgChildren.get("00");
                    if (bankBinField != null) {
                        gpayFtToBankRequest.setBankCode(bankBinField.value);
                    }

                    // Sub 01 - AccountNumber
                    EmvField accField = orgChildren.get("01");
                    if (accField != null) {
                        gpayFtToBankRequest.setAccountNumber(accField.value);
                    }
                } else {
                    // Fallback: tách 6 số đầu làm BIN nếu không parse được
                    String v = orgField.value;
                    if (v.length() > 6) {
                        gpayFtToBankRequest.setBankCode(v.substring(0, 6));
                        gpayFtToBankRequest.setAccountNumber(v.substring(6));
                    } else {
                        gpayFtToBankRequest.setAccountNumber(v);
                    }
                }
            }

        }

        // 54 - Amount
        EmvField f54 = rootMap.get("54");
        if (f54 != null) {
            gpayFtToBankRequest.setAmount(new BigDecimal(f54.value));
        }


        // 59 - Account / Merchant name
        EmvField f59 = rootMap.get("59");
        if (f59 != null) {
            gpayFtToBankRequest.setFullName(f59.value);
        }


        // 62 - Additional Data Field Template
        EmvField f62 = rootMap.get("62");
        if (f62 != null) {
            if (f62.children == null) {
                f62.children = parseNestedTemplate(f62.value);
            }
            Map<String, EmvField> addMap = f62.children;

            // 08 Purpose of desc
            if (addMap.get("08") != null) {
                gpayFtToBankRequest.setMessage(addMap.get("08").value);
            }

        }
        return gpayFtToBankRequest;
    }

    // Test nhanh
    public static void main(String[] args) {
        // Ví dụ minh hoạ, bạn thay bằng QR thật của bank để test
        String qr = "00020101021238520010A000000727012200069704250108454545450208QRIBFTTA530370454033005802VN62080804adam6304AFA9";
        Gson gson = new GsonBuilder()
                .setPrettyPrinting()   // in JSON đẹp, có format
                .serializeNulls()      // cho phép in cả giá trị null
                .create();

        GpayFtToBankRequest gpayFtToBankRequest = decodeVietQrToJson(qr);
        gpayFtToBankRequest.setType("ACCOUNT_NUMBER");
        System.out.println(gson.toJson(gpayFtToBankRequest));
    }

}
