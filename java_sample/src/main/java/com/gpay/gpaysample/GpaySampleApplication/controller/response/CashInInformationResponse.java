package com.gpay.gpaysample.GpaySampleApplication.controller.response;

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
public class CashInInformationResponse {
    private String accountNumber;
    private String accountName;
    private String bankCode;
    private String bankName;
    private String qrCode;
    private String qrCodeImage;

}
