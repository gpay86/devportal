package com.gpay.gpaysample.GpaySampleApplication.controller.webhook;

import com.gpay.gpaysample.GpaySampleApplication.controller.request.AccountBalanceWebhookRequest;
import com.gpay.gpaysample.GpaySampleApplication.controller.request.GpayPaymentGatewayCallbackRequest;
import com.gpay.gpaysample.GpaySampleApplication.exception.GpayBadRequestException;
import com.gpay.gpaysample.GpaySampleApplication.service.GpayWebhookService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;

@RestController
@RequiredArgsConstructor
public class GpayWebhookController {

    private final GpayWebhookService gpayWebhookService;

    /**
     * Controller này xử lý các thông báo được gửi lại từ GPay sau khi thanh toán thành công
     * (This controller handles information which are sent from Gpay after payment success  )
     *
     * @param gpayInternationalCallbackRequest
     * @return - log info của request từ Gpay trả về
     * (log info of request from Gpay callback)
     */
    @PostMapping("webhook/payment/international")
    public ResponseEntity<Object> internationalWebhook(@Valid @RequestBody GpayPaymentGatewayCallbackRequest gpayInternationalCallbackRequest) {
        return gpayWebhookService.verifyWebhookPaymentGateway(gpayInternationalCallbackRequest);
    }

    /**
     * Controller này để xử lý, hứng thông báo được gửi bới Gpay mỗi khi có sự thay đổi số dư trong tài khoảng ảo
     * (This controller handles notifications sent by Gpay whenever there are changes in the balance of a virtual account.)
     *
     * @param requestFromGpay - Request(notification) được gửi từ Gpay
     *                        (Request(notification) are sent from Gpay)
     * @return - log info của request của Gpay
     * (log info of request from Gpay)
     */
    @PostMapping("webhook/balance")
    public ResponseEntity<Object> handleBalanceUpdate(@RequestBody AccountBalanceWebhookRequest requestFromGpay) {
        //logic
        return gpayWebhookService.verifyBalanceUpdate(requestFromGpay);
    }

}
