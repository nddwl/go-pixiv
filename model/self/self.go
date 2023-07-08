package self

//Self https://www.pixiv.net/self?lang=<lang>&version=<version>
type Self struct {
	UserData struct {
		ID            string `json:"id"`
		PixivID       string `json:"pixivId"`
		Name          string `json:"name"`
		ProfileImg    string `json:"profileImg"`
		ProfileImgBig string `json:"profileImgBig"`
		Premium       bool   `json:"premium"`
		XRestrict     int    `json:"xRestrict"`
		Adult         bool   `json:"adult"`
		SafeMode      bool   `json:"safeMode"`
		IllustCreator bool   `json:"illustCreator"`
		NovelCreator  bool   `json:"novelCreator"`
		HideAiWorks   bool   `json:"hideAiWorks"`
	} `json:"userData"`
	Lang     string `json:"lang"`
	Token    string `json:"token"`
	Services struct {
		Booth    string `json:"booth"`
		Sketch   string `json:"sketch"`
		VroidHub string `json:"vroidHub"`
		Accounts string `json:"accounts"`
	} `json:"services"`
	Premium struct {
		FreeCampaign bool `json:"freeCampaign"`
	} `json:"premium"`
	Development bool `json:"development"`
	MiscData    struct {
		Consent struct {
			Gdpr bool `json:"gdpr"`
		} `json:"consent"`
		PolicyRevision bool `json:"policyRevision"`
		Grecaptcha     struct {
			RecaptchaEnterpriseScoreSiteKey string `json:"recaptchaEnterpriseScoreSiteKey"`
		} `json:"grecaptcha"`
		Info struct {
			ID         string `json:"id"`
			Title      string `json:"title"`
			CreateDate string `json:"createDate"`
		} `json:"info"`
		IsSmartphone   bool   `json:"isSmartphone"`
		OneSignalAppID string `json:"oneSignalAppId"`
	} `json:"miscData"`
	Mute []struct {
		Type        int    `json:"type"`        //屏蔽类型[0:"标签",1:"用户"]
		Value       string `json:"value"`       //Value 标签名称或作者名称
		PremiumSlot bool   `json:"premiumSlot"` //PremiumSlot 会员相关
	} `json:"mute"` //Mute 屏蔽列表,非会员最多屏蔽一个
	ActiveABTests struct {
		AdsApsEmailHash                            bool `json:"ads_aps_email_hash"`
		PlanAiTypePhase2                           bool `json:"plan_ai_type_phase_2"`
		DashboardNextJs                            bool `json:"dashboard_next_js"`
		AbIlllustSeriesSpaDev                      bool `json:"ab_illlust_series_spa_dev"`
		AbIlllustSeriesSpaNologinDev               bool `json:"ab_illlust_series_spa_nologin_dev"`
		IllustReplyTree                            bool `json:"illust_reply_tree"`
		AbMangaEachLangPopulerWorks                bool `json:"ab_manga_each_lang_populer_works"`
		AbMangaNewViewer                           bool `json:"ab_manga_new_viewer"`
		Nagisa                                     bool `json:"nagisa"`
		Novel12ThPremiumCovers                     bool `json:"novel_12th_premium_covers"`
		NovelMangaReadingStatusAlphaOnlyRequest    bool `json:"novel_manga_reading_status_alpha_only_request"`
		NovelMovingTopSection                      bool `json:"novel_moving_top_section"`
		TouchNovelFollowWatchlistTab               bool `json:"touch_novel_follow_watchlist_tab"`
		NovelCoverEditLineBreak                    bool `json:"novel_cover_edit_line_break"`
		PostedNovelCoverEdit                       bool `json:"posted_novel_cover_edit"`
		RecommendTutorial20191213                  bool `json:"recommend_tutorial_20191213"`
		RecommendNovelOnIllustDetailsPageC20230207 bool `json:"recommend_novel_on_illust_details_page_c_20230207"`
		SearchFilterAiType                         bool `json:"search_filter_ai_type"`
		TouchTopJack                               bool `json:"touch_top_jack"`
		WwwPremiumLinkText                         bool `json:"www_premium_link_text"`
		WwwTagsLinkToEnDic                         bool `json:"www_tags_link_to_en_dic"`
		WwwIllustUploadNextJs                      bool `json:"www_illust_upload_next_js"`
		WwwIllustReuploadNextJsDesktop             bool `json:"www_illust_reupload_next_js_desktop"`
	} `json:"activeABTests"`
	ActiveToggles struct {
		ToggleCommissionStopAiPhase1                 bool `json:"toggle_commission_stop_ai_phase_1"`
		ToggleCommissionAbilityToChangeResendRequest bool `json:"toggle_commission_ability_to_change_resend_request"`
		ToggleCommissionGuidelineNew                 bool `json:"toggle_commission_guideline_new"`
		ToggleMangaTutorialModal                     bool `json:"toggle_manga_tutorial_modal"`
		ToggleMangaNewViewerTutorialDb               bool `json:"toggle_manga_new_viewer_tutorial_db"`
		ToggleMangaThumbnailCrop                     bool `json:"toggle_manga_thumbnail_crop"`
		ToggleNovelMangaReadingStatusFrontendCache   bool `json:"toggle_novel_manga_reading_status_frontend_cache"`
		ToggleNovelWordCount                         bool `json:"toggle_novel_word_count"`
		ToggleFactoryCreateBySeries                  bool `json:"toggle_factory_create_by_series"`
		ToggleFactoryChangePageLimit                 bool `json:"toggle_factory_change_page_limit"`
	} `json:"activeToggles"`
}
