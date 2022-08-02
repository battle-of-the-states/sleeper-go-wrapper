package sleeper

// "github.com/PuerkitoBio/goquery"

// func (c *Client) GetAvatar(avatarID string) (string, error) {
// 	// https://sleepercdn.com/avatars/<avatar_id>
// 	lastfmURL := fmt.Sprintf("%s/avatars/%s", c.sleeperCDN, avatarID)
// 	fmt.Println(lastfmURL)

// 	resp, err := c.httpClient.Get(lastfmURL)

// 	if err != nil {
// 		return "", err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
// 	}

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	title := doc.Find("title").Text()

// 	return title, nil
// }
