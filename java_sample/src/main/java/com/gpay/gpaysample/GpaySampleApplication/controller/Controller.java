package com.gpay.gpaysample.GpaySampleApplication.controller;

import com.gpay.gpaysample.GpaySampleApplication.controller.request.FtToBankRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.PaymentGatewayInitRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.VaCreationRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.response.GeneralResponse;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayBadRequestException;
import com.gpay.gpaysample.GpaySampleApplication.service.GpayService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;

@RestController
@RequiredArgsConstructor
public class Controller {
    private final GpayService gpayService;


    /**
     * This is inquiry name banking by account number
     * @param request
     * @return
     * @throws GpayBadRequestException
     */
//    @PostMapping("ibft/inquiry")
//    public ResponseEntity<GeneralResponse> inquiry(@Valid @RequestBody BankAccountInfoRequest request) throws GpayBadRequestException {
//        return inquiryService.inquiry(request);
//    }

    /**
     * This is fund transfer to bank
     *
     * @param request
     * @return
     * @throws
     */
    @PostMapping("ibft/transfer")
    public ResponseEntity<GeneralResponse> ftToBank(@Valid @RequestBody FtToBankRequest request) {
        return gpayService.transferToBank(request);
    }


    @PostMapping("/va/create")
    public ResponseEntity<GeneralResponse> createVirtualAccount(@Valid @RequestBody VaCreationRequest request) {
        return gpayService.createVirtualAccount(request);
    }

    @PostMapping("/payment-gateway/init")
    public ResponseEntity<GeneralResponse> initPaymentGateway(@Valid @RequestBody PaymentGatewayInitRequest request) {
        return gpayService.initPaymentGateway(request);
    }

}
