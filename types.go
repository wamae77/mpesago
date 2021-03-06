package mpesa

const endpoint = "https://sandbox.safaricom.co.ke"

type (
	StkPushRequest struct {
		BusinessShortCode,
		Password,
		Timestamp,
		TransactionType,
		Amount,
		PhoneNumber,
		PartyA,
		PartyB,
		CallBackURL,
		QueueTimeOutURL,
		AccountReference,
		TransactionDesc string
	}

	AuthResponse struct {
		Access_token string
		Expires_in   string
	}

	Credetials struct {
		APP_KEY,
		APP_SECRET string
	}

	StkPushStatusRequest struct {
		Password          string
		Timestamp         string
		CheckoutRequestID string
		PhoneNumber       string
		BusinessShortCode string
	}
)
