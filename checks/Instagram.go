package mediaunlocktest

import (
	"io"
	"net/http"
	"strings"
)

func Instagram(c http.Client) Result {
	resp, err := PostForm(c, "https://www.instagram.com/api/graphql", `av=0&__d=www&__user=0&__a=1&__req=3&__hs=19750.HYP%3Ainstagram_web_pkg.2.1..0.0&dpr=1&__ccg=UNKNOWN&__rev=1011068636&__s=drshru%3Agu4p3s%3A0d8tzk&__hsi=7328972521009111950&__dyn=7xeUjG1mxu1syUbFp60DU98nwgU29zEdEc8co2qwJw5ux609vCwjE1xoswIwuo2awlU-cw5Mx62G3i1ywOwv89k2C1Fwc60AEC7U2czXwae4UaEW2G1NwwwNwKwHw8Xxm16wUwtEvw4JwJCwLyES1Twoob82ZwrUdUbGwmk1xwmo6O1FwlE6PhA6bxy4UjK5V8&__csr=gtneJ9lGF4HlRX-VHjmipBDGAhGuWV4uEyXyp22u6pU-mcx3BCGjHS-yabGq4rhoWBAAAKamtnBy8PJeUgUymlVF48AGGWxCiUC4E9HG78og01bZqx106Ag0clE0kVwdy0Nx4w2TU0iGDgChwmUrw2wVFQ9Bg3fw4uxfo2ow0asW&__comet_req=7&lsd=AVrkL73GMdk&jazoest=2909&__spin_r=1011068636&__spin_b=trunk&__spin_t=1706409389&fb_api_caller_class=RelayModern&fb_api_req_friendly_name=PolarisPostActionLoadPostQueryQuery&variables=%7B%22shortcode%22%3A%22C2YEAdOh9AB%22%2C%22fetch_comment_count%22%3A40%2C%22fetch_related_profile_media_count%22%3A3%2C%22parent_comment_count%22%3A24%2C%22child_comment_count%22%3A3%2C%22fetch_like_count%22%3A10%2C%22fetch_tagged_user_count%22%3Anull%2C%22fetch_preview_comment_count%22%3A2%2C%22has_threaded_comments%22%3Atrue%2C%22hoisted_comment_id%22%3Anull%2C%22hoisted_reply_id%22%3Anull%7D&server_timestamps=true&doc_id=10015901848480474`,
		H{"Accept", "*/*"},
		H{"Accept-Language", "zh-CN,zh;q=0.9"},
		H{"Connection", "keep-alive"},
		H{"Content-Type", "application/x-www-form-urlencoded"},
		H{"Cookie", "csrftoken=mmCtHhtfZRG-K3WgoYMemg; dpr=1.75; _js_ig_did=809EA442-22F7-4844-9470-ABC2AC4DE7AE; _js_datr=rb21ZbL7KR_5DN8m_43oEtgn; mid=ZbW9rgALAAECR590Ukv8bAlT8YQX; ig_did=809EA442-22F7-4844-9470-ABC2AC4DE7AE; ig_nrcb=1"},
		H{"Origin", "https://www.instagram.com"},
		H{"Referer", "https://www.instagram.com/p/C2YEAdOh9AB/"},
		H{"X-ASBD-ID", "129477"},
		H{"X-CSRFToken", "mmCtHhtfZRG-K3WgoYMemg"},
		H{"X-FB-Friendly-Name", "PolarisPostActionLoadPostQueryQuery"},
		H{"X-FB-LSD", "AVrkL73GMdk"},
		H{"X-IG-App-ID", "936619743392459"},
		H{"dpr", "1.75"},
		H{"sec-ch-prefers-color-scheme", "light"},
		H{"sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`},
		H{"sec-ch-ua-full-version-list", `"Not_A Brand";v="8.0.0.0", "Chromium";v="120.0.6099.225", "Google Chrome";v="120.0.6099.225"`},
		H{"sec-ch-ua-mobile", "?0"},
		H{"sec-ch-ua-model", `""`},
		H{"sec-ch-ua-platform", `"macOS"`},
		H{"sec-ch-ua-platform-version", `"11.14.6"`},
		H{"viewport-width", "1640"},
	)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if err != nil {
		return Result{Status: StatusFailed}
	}

	if resp.StatusCode == 200 {
		if strings.Contains(bodyString, `"should_mute_audio":true`) {
			return Result{Status: StatusNo}
		}
		return Result{Status: StatusOK}
	}

	return Result{Status: StatusUnexpected}
}
