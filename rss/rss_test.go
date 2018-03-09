package rss_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/yujiahaol68/rossy/rss"
)

const (
	rssContent = `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - Home]]></title>
        <description><![CDATA[BBC News - Home]]></description>
        <link>http://www.bbc.co.uk/news/</link>
        <image>
            <url>http://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - Home</title>
            <link>http://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Thu, 08 Mar 2018 09:38:37 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[Russian spy: Police seek to identify nerve agent source]]></title>
            <description><![CDATA[A police officer taken ill alongside an ex-Russian spy and his daughter is now awake and talking.]]></description>
            <link>http://www.bbc.co.uk/news/uk-43326734</link>
            <guid isPermaLink="true">http://www.bbc.co.uk/news/uk-43326734</guid>
            <pubDate>Thu, 08 Mar 2018 08:55:36 GMT</pubDate>
            <media:thumbnail width="976" height="549" url="http://c.files.bbci.co.uk/FB10/production/_100327246_mediaitem100327245.jpg"/>
        </item>
        <item>
            <title><![CDATA[CCTV shows poisoned spy buying scratchcards]]></title>
            <description><![CDATA[Footage has emerged of Sergei Skripal at a shop in Salisbury just five days before he collapsed.]]></description>
            <link>http://www.bbc.co.uk/news/world-43323506</link>
            <guid isPermaLink="true">http://www.bbc.co.uk/news/world-43323506</guid>
            <pubDate>Wed, 07 Mar 2018 17:27:47 GMT</pubDate>
            <media:thumbnail width="976" height="549" url="http://c.files.bbci.co.uk/A930/production/_100321334_p060dhbb.jpg"/>
        </item>
        <item>
            <title><![CDATA[Florida shooting: Gun control law moves step closer]]></title>
            <description><![CDATA[State lawmakers pass a bill that raises the legal age to buy a firearm from 18 to 21.]]></description>
            <link>http://www.bbc.co.uk/news/world-us-canada-43325913</link>
            <guid isPermaLink="true">http://www.bbc.co.uk/news/world-us-canada-43325913</guid>
            <pubDate>Thu, 08 Mar 2018 07:36:07 GMT</pubDate>
            <media:thumbnail width="976" height="549" url="http://c.files.bbci.co.uk/FCFA/production/_100326746_mediaitem100326745.jpg"/>
        </item>
        <item>
            <title><![CDATA[Kim Wall death: Inventor Peter Madsen goes on trial]]></title>
            <description><![CDATA[Peter Madsen faces charges of killing and dismembering Swedish journalist Kim Wall.]]></description>
            <link>http://www.bbc.co.uk/news/world-europe-43325462</link>
            <guid isPermaLink="true">http://www.bbc.co.uk/news/world-europe-43325462</guid>
            <pubDate>Thu, 08 Mar 2018 02:26:33 GMT</pubDate>
            <media:thumbnail width="976" height="549" url="http://c.files.bbci.co.uk/A9BB/production/_97315434_sub-composite.jpg"/>
        </item>
    </channel>
</rss>`
)

func Test_parse(t *testing.T) {
	r := rss.New()
	err := xml.Unmarshal([]byte(rssContent), r)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range r.ItemList {
		fmt.Printf("* %s\n%s\n", item.Title, item.Link)
	}
}
