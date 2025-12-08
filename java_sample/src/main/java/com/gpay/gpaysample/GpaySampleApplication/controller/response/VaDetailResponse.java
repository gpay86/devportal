package com.gpay.gpaysample.GpaySampleApplication.controller.response;

import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
public class VaDetailResponse {
    private String accountNumber;
    private String accountName;
    private String status;
    private String startAt;
    private String expireAt;
    private int balance;
    private String accountType;
    private Integer maxAmount;
    private Integer minAmount;
    private Integer equalAmount;
    private String qrCode;
    private String qrCodeImage;
    private String signature;
}
