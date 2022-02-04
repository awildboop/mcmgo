package mcmgo

const (
	gatewayUrl string = "https://api.mc-market.org/"

	apiHealth = gatewayUrl + "health/"
	apiAlerts = gatewayUrl + "alerts/"

	apiConversations        = "conversations/"
	apiConversationsReplies = apiConversations + "%s/replies" // %s = {c_id}

	apiMembers      = "members/%s/" // %s = {m_id}
	apiProfilePosts = "members/self/profile-posts/"

	apiResourcesBase = "resources/"
	apiResources     = apiResourcesBase + "%s/"           // %s = {r_id}
	apiVersions      = apiResourcesBase + "%s/versions/"  // %s = {r_id}
	apiUpdates       = apiResourcesBase + "%s/updates/"   // %s = {r_id}
	apiReviews       = apiResourcesBase + "%s/reviews/"   // %s = {r_id}
	apiPurchases     = apiResourcesBase + "%s/purchases/" // %s = {r_id}
	apiLicenses      = apiResourcesBase + "%s/licenses/"  // %s = {r_id}
	apiDownloads     = apiResourcesBase + "%s/downloads/" // %s = {r_id}
)

var (
	session *Session
)
