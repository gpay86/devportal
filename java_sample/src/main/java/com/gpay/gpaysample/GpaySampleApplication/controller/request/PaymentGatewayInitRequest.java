package com.gpay.gpaysample.GpaySampleApplication.controller.request;

import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
public class PaymentGatewayInitRequest {
    private String address;
    private int amount;
    private String callbackUrl;
    private String customerId;
    private String customerName;
    private String description;
    private String email;
    private String embedData;
    private String paymentMethod;
    private String paymentType;
    private String phone;
    private String requestId;
    private String title;
    private String webhookUrl;
}
