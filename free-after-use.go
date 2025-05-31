const MAX_RECEIPT_ID_LENGTH = 8
const RECEIPT_EXPIRATION_TIME = 1800 * time.Second

func externalHandler(w http.ResponseWriter, r *http.Request) {
    receiptIdInt := ""
    vars := mux.Vars(r)
    receiptId := vars["receiptId"]
    sessionId, _ := r.Cookie("sessionId")

    if len(receiptId) > 8 {
        return
    }

    regexInt := regexp.MustCompile(`[0-9]+`)
    match := regexInt.FindStringSubmatch(receiptId)
    if len(match) != 0 {
        receiptIdInt = match[0]
    } else {
        return
    }

    if !isAuthenticated(sessionId.Value) {
        return
    }

    if isAuthorized(receiptId, sessionId.Value) {
        releaseReceipt(receiptId, 1800*time.Second)
    }

    if isReceiptReleased(receiptId) {
        billingURL := strings.Replace("https://internal-host/api/v1/billing/receipt/:receiptId", ":receiptId", receiptIdInt, 1)

        req, _ := http.NewRequest("GET", billingURL, nil)
        req.Header.Set("Authorization", getToken())
        client := &http.Client{}
        resp, _ := client.Do(req)

        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)
        w.Write(body)
    }
}
