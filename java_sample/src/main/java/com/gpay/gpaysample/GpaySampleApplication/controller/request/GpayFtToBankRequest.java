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
public class GpayFtToBankRequest {
    private String accountNumber;
    private String bankCode;
    private String fullName;
    private Integer amount;
    private String transactionId;
    private String type;
    private String orderRef;
    private String mapId;
    private String message;

}
