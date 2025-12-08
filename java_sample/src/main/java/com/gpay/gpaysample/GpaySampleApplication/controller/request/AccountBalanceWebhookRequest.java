package com.gpay.gpaysample.GpaySampleApplication.controller.request;

import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
public class AccountBalanceWebhookRequest {
    private String gpayTransId;
    private String bankTraceId;
    private String bankTransactionId;
    private String accountNumber;
    private Integer amount;
    private String message;
    private String merchantCode;
    private String action;
    private String signature;

}
