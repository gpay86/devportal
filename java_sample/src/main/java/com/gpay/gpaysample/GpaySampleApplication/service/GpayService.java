package com.gpay.gpaysample.GpaySampleApplication.service;

import com.gpay.gpaysample.GpaySampleApplication.controller.request.FtToBankRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.GpayFtToBankRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.PaymentGatewayInitRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.VaCreationRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.response.GeneralResponse;
import com.gpay.gpaysample.GpaySampleApplication.enumeration.IbftApiUri;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayBadRequestException;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import java.util.Objects;

@Service
@RequiredArgsConstructor
public class GpayService extends BaseService {

    public static final String TRANSFER_PATH = "/payouts/instant/transfer-to-bank";
    public static final String CREATE_VA_PATH = "/collection/va/create";
    public static final String PAYMENT_GATEWAY_INIT_PATH = "/payments/gateway/init-order";

    public ResponseEntity<GeneralResponse> transferToBank(FtToBankRequest ftToBankRequest) {
        GpayFtToBankRequest gpayFtToBankRequest = new GpayFtToBankRequest();
        gpayFtToBankRequest.setAmount(ftToBankRequest.getAmount());
        gpayFtToBankRequest.setTransactionId(ftToBankRequest.getTransactionId());
        gpayFtToBankRequest.setAccountNumber(ftToBankRequest.getAccountNumber());
        gpayFtToBankRequest.setBankCode(ftToBankRequest.getBankCode());
        gpayFtToBankRequest.setFullName(ftToBankRequest.getFullName());
        gpayFtToBankRequest.setType(ftToBankRequest.getType());
        gpayFtToBankRequest.setOrderRef(ftToBankRequest.getOrderRef());
        gpayFtToBankRequest.setMapId(ftToBankRequest.getMapId());
        gpayFtToBankRequest.setMessage(ftToBankRequest.getMessage());
        GeneralResponse response = baseRestTemplate.request(TRANSFER_PATH, gpayFtToBankRequest);
        return ResponseEntity.ok(response);
    }

    public ResponseEntity<GeneralResponse> createVirtualAccount(VaCreationRequest request) {
        VaCreationRequest vaCreationRequest = new VaCreationRequest();
        vaCreationRequest.setAccountType(request.getAccountType());
        vaCreationRequest.setBankCode(request.getBankCode());
        vaCreationRequest.setAccountName(request.getAccountName());
        vaCreationRequest.setMapType(request.getMapType());
        vaCreationRequest.setMapId(request.getMapId());
        vaCreationRequest.setEqualAmount(request.getEqualAmount());
        vaCreationRequest.setMinAmount(request.getMinAmount());
        vaCreationRequest.setMaxAmount(request.getMaxAmount());
        GeneralResponse response = baseRestTemplate.request(CREATE_VA_PATH, vaCreationRequest);
        return ResponseEntity.ok(response);
    }

    public ResponseEntity<GeneralResponse> initPaymentGateway(PaymentGatewayInitRequest request) {

        GeneralResponse response = baseRestTemplate.request(PAYMENT_GATEWAY_INIT_PATH, request);
        return ResponseEntity.ok(response);
    }

}
