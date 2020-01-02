package private

type API struct {
	KEY    string
	SECRET string
}

func (api *API) sign() string {

	return ""
}

func (api *API) send() error {

	return nil
}

/*
PHP EXAMPLE FROM DEV API DOC:
<?php
  $bitsoKey = "API_KEY";
  $bitsoSecret = "API_SECRET"
  $nonce = round(microtime(true) * 1000);
  $HTTPMethod = "POST";
  $RequestPath = "/v3/orders/";
  $JSONPayload = json_encode(['book'  => 'btc_mxn',
                              'side'  => 'buy',
                              'major' => '.01',
                              'price' => '1000',
                              'type'  => 'limit']);

  // Create signature
  $message = $nonce . $HTTPMethod . $RequestPath . $JSONPayload;
  $signature = hash_hmac('sha256', $message, $bitsoSecret);

  // Build the auth header
  $format = 'Bitso %s:%s:%s';
  $authHeader =  sprintf($format, $bitsoKey, $nonce, $signature);


  // Send request
  $ch = curl_init();
  curl_setopt($ch, CURLOPT_URL, 'https://api.bitso.com/v3/orders/');
  curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");
  curl_setopt($ch, CURLOPT_POSTFIELDS, $JSONPayload);
  curl_setopt($ch, CURLOPT_HTTPHEADER, array(
      'Authorization: ' .  $authHeader,
      'Content-Type: application/json'));
  $result = curl_exec($ch);

  echo $result;
?>
*/
