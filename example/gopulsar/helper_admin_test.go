package gopulsar

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

// TODO: add utils to manipulate admins
// https://github.com/streamnative/pulsarctl
// 1. DeleteAllTopic
// 2. DeleteOneTopic
// 3. DeleteTopics(pattern)

func defaultNS() utils.NameSpaceName {
	ns, err := utils.GetNameSpaceName("public", "default")
	if err != nil {
		panic(err)
	}
	return *ns
}

func newAdminClient() pulsar.Client {
	config := &common.Config{
		WebServiceURL:              "http://localhost:8080",
		TLSAllowInsecureConnection: true,
		PulsarAPIVersion:           common.V2,
	}

	admin, err := pulsar.New(config)
	if err != nil {
		panic(err)
	}

	return admin
}

func TestPulsarctl_Topics(t *testing.T) {
	admin := newAdminClient()

	// Partitioned Topics, NonPartitioned Topics
	pTopics, npTopics, err := admin.Topics().List(defaultNS())
	fmt.Println(pTopics)
	fmt.Println(npTopics)
	fmt.Println(err)

}

type Topics []string

func (t Topics) String() string {
	switch len(t) {
	case 0:
		return ""
	case 1:
		return t[0]
	default:
		return "\n" + strings.Join(t, "\n") + "\n"
	}
}

func getAllTopics() Topics {
	admin := newAdminClient()

	// Partitioned Topics, NonPartitioned Topics
	// For standalone mode, only non-partitioned topics are available.
	_, npTopics, err := admin.Topics().List(defaultNS())
	if err != nil {
		panic(err)
	}

	return Topics(npTopics)
}

func deleteTopic(pattern string) {
	re := regexp.MustCompile(pattern)
	admin := newAdminClient()
	doDelete := func(completeName string) error {
		topicName, err := utils.GetTopicName(completeName)
		if err != nil {
			return err
		}

		err = admin.Topics().Delete(*topicName, true, true)
		return err
	}

	// For standalone mode, only non-partitioned topics are available.
	_, npTopics, err := admin.Topics().List(defaultNS())
	if err != nil {
		panic(err)
	}

	if len(npTopics) == 0 {
		return
	}

	for _, t := range npTopics {
		if !re.MatchString(t) {
			continue
		}
		fmt.Printf("===> deleting topic: %v\n", t)
		if err := doDelete(t); err != nil {
			panic(err)
		}
	}
}

func disableCompact() {
	admin := newAdminClient()
	ns := defaultNS()
	err := admin.Namespaces().SetCompactionThreshold(ns, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("disabled compaction on ns: %v\n", ns)
}

func getTopicCompleteName(topic string) string {
	return "persistent://public/default/" + topic
}

func deleteSubscription(topic, subName string) {
	admin := newAdminClient()
	topicName := getTopicCompleteName(topic)
	tn, err := utils.GetTopicName(topicName)

	if err != nil {
		panic(err)
	}

	err = admin.Subscriptions().Delete(*tn, subName)
	if err != nil {
		panic(err)
	}
}

func TestDeleteTopic_Pattern(t *testing.T) {
	topicsBefore := getAllTopics()
	fmt.Printf("==> topics before: [%+v]\n\n", topicsBefore)

	// pattern := ".*topic-ack-multiple.*"
	// pattern := ".*topic-\\d.*"
	pattern := "topic-clj"
	deleteTopic(pattern)

	topicsAfter := getAllTopics()
	fmt.Printf("\n==> topics after: [%+v]\n\n", topicsAfter)
}

func TestDeleteTopic_All(t *testing.T) {
	deleteTopic(".*")

	topicsAfter := getAllTopics()
	fmt.Printf("\n==> topics after: [%+v]\n\n", topicsAfter)
}

func TestCompact_Disable(t *testing.T) {
	disableCompact()
}

func TestTopic_ListAll(t *testing.T) {
	ts := getAllTopics()

	fmt.Printf("topics: \n%v\n", ts)

}
