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
public class VaCreationRequest {
    private String merchantCode;
    private String accountName;
    private String mapId;
    private String mapType;
    private String accountType;
    private String bankCode;
    private Integer maxAmount;
    private Integer minAmount;
    private Integer equalAmount;
    private String customerAddress;
}
