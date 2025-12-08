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
public class MerchantAccountInfoResponse {
    private Integer amountRevenue;
    private Integer amountCash;
    private Integer amountMinimum;
    private Integer amountRefund;
    private CashInInformationResponse cashInInformation;
}
