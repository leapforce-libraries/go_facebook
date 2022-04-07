package facebook

import (
	errortools "github.com/leapforce-libraries/go_errortools"
)

type IgStoryInsightsMetric IgMediaInsightsMetric

const (
	IgStoryInsightsMetricExits       IgStoryInsightsMetric = "exits"
	IgStoryInsightsMetricImpressions IgStoryInsightsMetric = "impressions"
	IgStoryInsightsMetricReach       IgStoryInsightsMetric = "reach"
	IgStoryInsightsMetricReplies     IgStoryInsightsMetric = "replies"
	IgStoryInsightsMetricTapsForward IgStoryInsightsMetric = "taps_forward"
	IgStoryInsightsMetricTapsBack    IgStoryInsightsMetric = "taps_back"
)

type GetIgStoryInsightsConfig struct {
	MediaId         string
	UserAccessToken *string
	Metrics         []IgStoryInsightsMetric
}

func (service *Service) GetIgStoryInsights(config *GetIgStoryInsightsConfig) (*[]IgMediaInsight, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetIgStoryInsightsConfig must not be a nil pointer")
	}

	metrics := []IgMediaInsightsMetric{}
	for _, metric := range config.Metrics {
		metrics = append(metrics, IgMediaInsightsMetric(metric))
	}

	getIgMediaInsightsConfig := GetIgMediaInsightsConfig{
		MediaId:         config.MediaId,
		UserAccessToken: config.UserAccessToken,
		Metrics:         metrics,
	}

	return service.GetIgMediaInsights(&getIgMediaInsightsConfig)
}
