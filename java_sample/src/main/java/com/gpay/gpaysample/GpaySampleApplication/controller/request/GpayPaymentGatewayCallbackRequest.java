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
public class GpayPaymentGatewayCallbackRequest {
    private String merchantOrderId;
    private String gpayTransId;
    private String gpayBillId;
    private String status;
    private String embedData;
    private String userPaymentMethod;
    private String signature;
}
