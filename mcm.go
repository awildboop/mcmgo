package mcmgo

const (
	gatewayUrl string = "https://api.mc-market.org/"

	apiHealth = gatewayUrl + "v1/health/"
	apiAlerts = gatewayUrl + "v1/alerts/"

	apiConversations        = gatewayUrl + "v1/conversations/"
	apiConversationsReplies = apiConversations + "%s/replies" // %s = {c_id}

	apiMembers      = gatewayUrl + "v1/members/%s/" // %s = {m_id}
	apiProfilePosts = gatewayUrl + "v1/members/self/profile-posts/"

	apiResourceBase = gatewayUrl + "/v1/resources/"
	apiResources    = apiResourceBase + "%s/"           // %s = {r_id}
	apiVersions     = apiResourceBase + "%s/versions/"  // %s = {r_id}
	apiUpdates      = apiResourceBase + "%s/updates/"   // %s = {r_id}
	apiReviews      = apiResourceBase + "%s/reviews/"   // %s = {r_id}
	apiPurchases    = apiResourceBase + "%s/purchases/" // %s = {r_id}
	apiLicenses     = apiResourceBase + "%s/licenses/"  // %s = {r_id}
	apDownloads     = apiResourceBase + "%s/downloads/" // %s = {r_id}
)

var ()
