package com.gpay.gpaysample.GpaySampleApplication.enumeration;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum IbftApiUri {
    INQUIRY("/fund-transfers/inquiry"),
    TRANFER("/fund-transfers/ft-to-bank"),
    DETAILEDTRANSACTION("/fund-transfers/"),
    MERCHANTBALANCE("/merchants/get-merchant-account-information");
    private final String uri;
}

