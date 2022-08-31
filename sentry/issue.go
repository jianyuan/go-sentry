package sentry

type IssueService service

type Issue struct {
	Activity            []Activity           `json:"activity"`
	Annotations         []interface{}        `json:"annotations"`
	AssignedTo          interface{}          `json:"assignedTo"`
	Count               string               `json:"count"`
	Culprit             string               `json:"culprit"`
	FirstRelease        FirstRelease         `json:"firstRelease"`
	FirstSeen           string               `json:"firstSeen"`
	HasSeen             bool                 `json:"hasSeen"`
	ID                  string               `json:"id"`
	IsBookmarked        bool                 `json:"isBookmarked"`
	IsPublic            bool                 `json:"isPublic"`
	IsSubscribed        bool                 `json:"isSubscribed"`
	LastRelease         interface{}          `json:"lastRelease"`
	LastSeen            string               `json:"lastSeen"`
	Level               string               `json:"level"`
	Logger              interface{}          `json:"logger"`
	Metadata            Metadata             `json:"metadata"`
	NumComments         int64                `json:"numComments"`
	Participants        []interface{}        `json:"participants"`
	Permalink           string               `json:"permalink"`
	PluginActions       []interface{}        `json:"pluginActions"`
	PluginContexts      []interface{}        `json:"pluginContexts"`
	PluginIssues        []interface{}        `json:"pluginIssues"`
	Project             WelcomeProject       `json:"project"`
	SeenBy              []interface{}        `json:"seenBy"`
	ShareID             interface{}          `json:"shareId"`
	ShortID             string               `json:"shortId"`
	Stats               map[string][][]int64 `json:"stats"`
	Status              string               `json:"status"`
	StatusDetails       interface{}          `json:"statusDetails"`
	SubscriptionDetails interface{}          `json:"subscriptionDetails"`
	Tags                []interface{}        `json:"tags"`
	Title               string               `json:"title"`
	Type                string               `json:"type"`
	UserCount           int64                `json:"userCount"`
	UserReportCount     int64                `json:"userReportCount"`
}

type Activity struct {
	Data        interface{} `json:"data"`
	DateCreated string      `json:"dateCreated"`
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	User        interface{} `json:"user"`
}

type FirstRelease struct {
	Authors      []interface{}    `json:"authors"`
	CommitCount  int64            `json:"commitCount"`
	Data         interface{}      `json:"data"`
	DateCreated  string           `json:"dateCreated"`
	DateReleased interface{}      `json:"dateReleased"`
	DeployCount  int64            `json:"deployCount"`
	FirstEvent   string           `json:"firstEvent"`
	LastCommit   interface{}      `json:"lastCommit"`
	LastDeploy   interface{}      `json:"lastDeploy"`
	LastEvent    string           `json:"lastEvent"`
	NewGroups    int64            `json:"newGroups"`
	Owner        interface{}      `json:"owner"`
	Projects     []ProjectElement `json:"projects"`
	Ref          interface{}      `json:"ref"`
	ShortVersion string           `json:"shortVersion"`
	URL          interface{}      `json:"url"`
	Version      string           `json:"version"`
}

type ProjectElement struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Metadata struct {
	Title string `json:"title"`
}

type WelcomeProject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
