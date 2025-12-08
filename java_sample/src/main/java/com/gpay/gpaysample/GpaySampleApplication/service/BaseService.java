package com.gpay.gpaysample.GpaySampleApplication.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.gpay.gpaysample.GpaySampleApplication.clients.BaseRestTemplate;
import com.gpay.gpaysample.GpaySampleApplication.configs.Config;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayBadRequestException;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayErrorCode;
import com.gpay.gpaysample.GpaySampleApplication.utils.CryptoUtils;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.ClassPathResource;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class BaseService {

    @Autowired
    protected Config config;
    @Autowired
    protected BaseRestTemplate baseRestTemplate;
    @Autowired
    protected ObjectMapper objectMapper;




}