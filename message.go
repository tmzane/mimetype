package mimetype

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

const (
	// MessageBhttp bhttp mime type.
	MessageBhttp = "message/bhttp"
	// MessageCpim CPIM mime type.
	MessageCpim = "message/CPIM"
	// MessageDeliveryStatus delivery-status mime type.
	MessageDeliveryStatus = "message/delivery-status"
	// MessageDispositionNotification disposition-notification mime type.
	MessageDispositionNotification = "message/disposition-notification"
	// MessageExample example mime type.
	MessageExample = "message/example"
	// MessageExternalBody external-body mime type.
	MessageExternalBody = "message/external-body"
	// MessageFeedbackReport feedback-report mime type.
	MessageFeedbackReport = "message/feedback-report"
	// MessageGlobal global mime type.
	MessageGlobal = "message/global"
	// MessageGlobalDeliveryStatus global-delivery-status mime type.
	MessageGlobalDeliveryStatus = "message/global-delivery-status"
	// MessageGlobalDispositionNotification global-disposition-notification mime type.
	MessageGlobalDispositionNotification = "message/global-disposition-notification"
	// MessageGlobalHeaders global-headers mime type.
	MessageGlobalHeaders = "message/global-headers"
	// MessageHTTP http mime type.
	MessageHTTP = "message/http"
	// MessageImdnXML imdn+xml mime type.
	MessageImdnXML = "message/imdn+xml"
	// MessageNewsObsoletedByRfc5537 news (OBSOLETED by [RFC5537]) mime type.
	MessageNewsObsoletedByRfc5537 = "message/news"
	// MessagePartial partial mime type.
	MessagePartial = "message/partial"
	// MessageRfc822 rfc822 mime type.
	MessageRfc822 = "message/rfc822"
	// MessageSHTTPObsolete s-http (OBSOLETE) mime type.
	MessageSHTTPObsolete = "message/s-http"
	// MessageSip sip mime type.
	MessageSip = "message/sip"
	// MessageSipfrag sipfrag mime type.
	MessageSipfrag = "message/sipfrag"
	// MessageTrackingStatus tracking-status mime type.
	MessageTrackingStatus = "message/tracking-status"
	// MessageVndSiSimpObsoletedByRequest vnd.si.simp (OBSOLETED by request) mime type.
	MessageVndSiSimpObsoletedByRequest = "message/vnd.si.simp"
	// MessageVndWfaWsc vnd.wfa.wsc mime type.
	MessageVndWfaWsc = "message/vnd.wfa.wsc"
)

// IsMessage checks if the mime types is message.
func IsMessage(mt string) bool {
	switch mt {
	case "message/CPIM":
		return true
	case "message/bhttp":
		return true
	case "message/delivery-status":
		return true
	case "message/disposition-notification":
		return true
	case "message/example":
		return true
	case "message/external-body":
		return true
	case "message/feedback-report":
		return true
	case "message/global":
		return true
	case "message/global-delivery-status":
		return true
	case "message/global-disposition-notification":
		return true
	case "message/global-headers":
		return true
	case "message/http":
		return true
	case "message/imdn+xml":
		return true
	case "message/news":
		return true
	case "message/partial":
		return true
	case "message/rfc822":
		return true
	case "message/s-http":
		return true
	case "message/sip":
		return true
	case "message/sipfrag":
		return true
	case "message/tracking-status":
		return true
	case "message/vnd.si.simp":
		return true
	case "message/vnd.wfa.wsc":
		return true
	default:
		return false
	}
}
