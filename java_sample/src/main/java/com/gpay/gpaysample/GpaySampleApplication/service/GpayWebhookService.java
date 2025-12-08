package com.gpay.gpaysample.GpaySampleApplication.service;

import com.google.gson.Gson;
import com.gpay.gpaysample.GpaySampleApplication.configs.Config;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.AccountBalanceWebhookRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.GpayPaymentGatewayCallbackRequest;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayBadRequestException;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayErrorCode;
import com.gpay.gpaysample.GpaySampleApplication.utils.CryptoUtils;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.apache.logging.log4j.util.Strings;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.ClassPathResource;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
@Log4j2
public class GpayWebhookService {

    final private Config config;
    final private Gson gson;

    public ResponseEntity verifyWebhookPaymentGateway(GpayPaymentGatewayCallbackRequest response) {
        log.info("Gpay payment gateway webhook {} ", gson.toJson(response));
        if (Strings.isNotBlank(response.getSignature())) {
            String stringData = "merchant_order_id=%s&gpay_trans_id=%s&gpay_bill_id=%s&status=%s&embed_data=%s&user_payment_method=%s";
            String rawData  = String.format(stringData, response.getMerchantOrderId(), response.getGpayTransId(), response.getGpayBillId(),
                    response.getStatus(), response.getEmbedData(), response.getUserPaymentMethod());
            if (verifySignature(rawData, response.getSignature())) {
                return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Unauthorized ");
            }
        }
        // TODO implement your business here
        return ResponseEntity.ok("OK");
    }

    public ResponseEntity verifyBalanceUpdate(AccountBalanceWebhookRequest response) {
        log.info("Gpay update balance webhook {}", gson.toJson(response));
        if (Strings.isNotBlank(response.getSignature())) {
            String stringData = "gpay_trans_id=%s&bank_trace_id=%s&bank_transaction_id=%s&account_number=%s&amount=%s&message=%s&action=%s";
            String rawData  = String.format(stringData, response.getGpayTransId(), response.getBankTraceId(), response.getBankTransactionId(),
                    response.getAccountNumber(), response.getAmount(),response.getMessage(), response.getAction());
            if (verifySignature(rawData, response.getSignature())) {
                return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Unauthorized ");
            }
        }
        // TODO implement your business here
        return ResponseEntity.ok("OK");
    }

    private boolean verifySignature(String dataString, String sig) {
        boolean signature;
        try {
            String publicKeyUrl = new ClassPathResource(config.getGpayPublicKey()).getFile().getAbsolutePath();
            signature = CryptoUtils.verifySignature(publicKeyUrl, dataString, sig);
        } catch (Exception e) {
            e.printStackTrace();
            return false;
        }
        return signature;
    }

}
