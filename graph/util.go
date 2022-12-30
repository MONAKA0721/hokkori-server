package graph

import (
	"context"
	"github.com/MONAKA0721/hokkori/ent"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
)

func GetNewHashtagIDs(ctx context.Context, client *ent.Client, hashtagTitles []*string) ([]int, error) {
	newHashtagIDs := make([]int, len(hashtagTitles))
	for index, hashtagTitle := range hashtagTitles {
		createdHashtag, err := client.Hashtag.Create().SetTitle(*hashtagTitle).Save(ctx)
		if err != nil {
			if ent.IsConstraintError(err) {
				queryHashtag, err := client.Hashtag.Query().Where(hashtag.TitleEQ(*hashtagTitle)).Only(ctx)
				if err != nil {
					return nil, err
				}
				newHashtagIDs[index] = queryHashtag.ID
			} else {
				return nil, err
			}
		} else {
			newHashtagIDs[index] = createdHashtag.ID
		}
	}
	return newHashtagIDs, nil
}
