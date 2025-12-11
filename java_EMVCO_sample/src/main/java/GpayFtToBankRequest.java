import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.math.BigDecimal;

@Getter
@Setter
@NoArgsConstructor
//@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
public class GpayFtToBankRequest {
    private String accountNumber;
    private String bankCode;
    private String fullName;
    private BigDecimal amount;
    private String transactionId;
    private String type;
    private String mapId;
    private String message;


}