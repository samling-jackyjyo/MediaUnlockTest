package mediaunlocktest

import (
    "io"
	"net/http"
	"strings"
)

func NeonTV(c http.Client) Result {
	resp, err := PostJson(c, "https://api.neontv.co.nz/api/client/gql?", `{"operationName":"UpdateAccount","variables":{"input":{"name":"Reid","surname":"Hershel","email":"restriction.check@gmail.com","password":"restriction.check","optIns":[{"id":"RECEIVE_UPDATES","subscribed":false}]}},"query":"mutation UpdateAccount($input: AccountInput!, $pin: String) {\n  account(input: $input, pin: $pin) {\n    ...AccountFields\n    __typename\n  }\n}\n\nfragment AccountFields on Account {\n  name\n  surname\n  email\n  selectedProfile\n  hasPin\n  optIns {\n    id\n    text\n    subscribed\n    __typename\n  }\n  phoneNumbers {\n    home\n    mobile\n    __typename\n  }\n  session {\n    token\n    __typename\n  }\n  profiles {\n    ...ProfileFields\n    __typename\n  }\n  settings {\n    requirePinForSwitchProfile\n    requirePinForManageProfile\n    tvodPurchaseRestriction\n    playbackQuality {\n      ...PlaybackQualityFields\n      __typename\n    }\n    __typename\n  }\n  purchases {\n    totalItems\n    items {\n      ...PurchaseFields\n      __typename\n    }\n    __typename\n  }\n  cpCustomerID\n  subscription {\n    ...SubscriptionInformationFields\n    __typename\n  }\n  __typename\n}\n\nfragment ProfileFields on Profile {\n  id\n  name\n  email\n  isKid\n  isDefault\n  needToConfigure\n  ageGroup\n  avatar {\n    uri\n    id\n    __typename\n  }\n  closedCaption\n  maxRating\n  mobile\n  __typename\n}\n\nfragment PlaybackQualityFields on PlaybackQuality {\n  wifi {\n    id\n    label\n    description\n    bitrate\n    __typename\n  }\n  __typename\n}\n\nfragment PurchaseFields on Purchase {\n  id\n  profile {\n    id\n    name\n    __typename\n  }\n  contentItem {\n    ...ContentItemLightFields\n    __typename\n  }\n  product {\n    id\n    name\n    renewable\n    __typename\n  }\n  total\n  startAvailable\n  endAvailable\n  endViewable\n  __typename\n}\n\nfragment ContentItemLightFields on ContentItem {\n  id\n  isRental\n  ... on Title {\n    id\n    ldId\n    path\n    title\n    year\n    rating {\n      id\n      rating\n      __typename\n    }\n    genres\n    duration\n    images {\n      uri\n      __typename\n    }\n    createdAt\n    products {\n      id\n      originalPrice\n      currentPrice\n      name\n      currency\n      __typename\n    }\n    isComingSoon\n    videoExtras {\n      ...VideoExtraFields\n      __typename\n    }\n    tile {\n      image\n      header\n      subHeader\n      badge\n      contentItem {\n        id\n        __typename\n      }\n      sortValues {\n        key\n        value\n        __typename\n      }\n      playbackInfo {\n        status\n        numberMinutesRemaining\n        numberMinutesWatched\n        position\n        __typename\n      }\n      rentalInfo {\n        secondsLeftToStartWatching\n        secondsLeftToWatch\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  ... on Series {\n    title\n    ldId\n    genres\n    path\n    products {\n      id\n      originalPrice\n      currentPrice\n      name\n      currency\n      __typename\n    }\n    seasons {\n      id\n      episodes {\n        id\n        title\n        seasonNumber\n        episodeNumber\n        __typename\n      }\n      __typename\n    }\n    images {\n      uri\n      __typename\n    }\n    createdAt\n    isComingSoon\n    videoExtras {\n      ...VideoExtraFields\n      __typename\n    }\n    tile {\n      image\n      header\n      subHeader\n      badge\n      contentItem {\n        id\n        __typename\n      }\n      sortValues {\n        key\n        value\n        __typename\n      }\n      playbackInfo {\n        status\n        numberMinutesRemaining\n        numberMinutesWatched\n        position\n        __typename\n      }\n      rentalInfo {\n        secondsLeftToStartWatching\n        secondsLeftToWatch\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  ... on Episode {\n    episodeNumber\n    seasonNumber\n    series {\n      id\n      title\n      path\n      seasons {\n        episodes {\n          id\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  ... on VideoExtra {\n    contentItems {\n      id\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment VideoExtraFields on VideoExtra {\n  id\n  description\n  images {\n    id\n    uri\n    height\n    width\n    __typename\n  }\n  tile {\n    image\n    __typename\n  }\n  start\n  end\n  title\n  videoEncodings {\n    ...VideoEncodingFields\n    __typename\n  }\n  __typename\n}\n\nfragment VideoEncodingFields on VideoEncoding {\n  id\n  format\n  referenceId\n  size\n  offlineEnabled\n  __typename\n}\n\nfragment SubscriptionInformationFields on SubscriptionInformation {\n  currentSubscription {\n    name\n    sku\n    endsAt\n    startsAt\n    price\n    features\n    order {\n      voucherCode\n      __typename\n    }\n    subscriptionGAType\n    promotion {\n      name\n      price\n      isSpark\n      isFreeTrial\n      expiration\n      isBridgingOfferPromotion\n      __typename\n    }\n    evSubscriptionStatus\n    __typename\n  }\n  upcomingSubscription {\n    name\n    sku\n    endsAt\n    startsAt\n    price\n    order {\n      voucherCode\n      __typename\n    }\n    subscriptionGAType\n    promotion {\n      name\n      price\n      isSpark\n      isFreeTrial\n      expiration\n      __typename\n    }\n    evSubscriptionStatus\n    __typename\n  }\n  upcomingFinalBillSubscription {\n    sku\n    evSubscriptionStatus\n    __typename\n  }\n  nextPayment {\n    date\n    method\n    type\n    price\n    __typename\n  }\n  recentPayments {\n    date\n    method\n    type\n    price\n    __typename\n  }\n  status\n  renewalStatus\n  recurringVouchers {\n    orderDetails {\n      productName\n      voucherCode\n      status\n      promotion {\n        endDate\n        id\n        amount\n        isExhausted\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  dcbSubscriptionInfo {\n    partnerName\n    __typename\n  }\n  __typename\n}\n"}`,
	    H{"authorization", "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiOTNlZDQyZC00M2RiLTQzNDMtYThjZi1mZTk4YTY2NzVkNTgiLCJpc3MiOiJMaWdodGJveCIsImV4cCI6MTY2NDk3NTExMSwiZGV2aWNlSWQiOiI2NTgzZWU4YmM0YzQwZmJhOTgzMGQ0ZTYwNTM5ZDAzNSIsInBsYXRmb3JtIjoiV2ViIiwiYnJvd3NlciI6IkNocm9tZSIsInRhYiI6MTYzMzQxNzUwNzY4OSwib3MiOiJXaW5kb3dzIDEwLjAiLCJpYXQiOjE2MzM0MTc1MTF9.E7qgVpqsJEPsh0B3lgK9x8hPs7nQ_Bio_FCt1H8mB4XCPrsand4kHVHA5LpiB5rvBLfOaSfJKru-gKuMlgLJhg"},
	)
	if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    
    if err != nil {
		return Result{Status: StatusNetworkErr}
	}
	
	if resp.StatusCode == 403 || strings.Contains(bodyString, "RESTRICTED_GEOLOCATION") {
		return Result{Status: StatusNo}
	}
	
	if resp.StatusCode == 200 || resp.StatusCode == 302 {
		return Result{Status: StatusOK}
	}
	
	return Result{Status: StatusUnexpected}
}