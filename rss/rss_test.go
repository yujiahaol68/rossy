package rss_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	"github.com/yujiahaol68/rossy/rss"
	"golang.org/x/net/html/charset"
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
    </channel>
</rss>`

	notUTF8rss = `<?xml version="1.0" encoding="gb2312"?>
    <?xml-stylesheet type="text/xsl" href="/css/rss_xml_style.css"?>
    <rss version="2.0">
      <channel>
        <title>新闻国内</title>
        <image>
          <title>新闻国内</title>
          <link>http://news.qq.com</link>
          <url>http://mat1.qq.com/news/rss/logo_news.gif</url>
        </image>
        <description>新闻国内</description>
        <link>http://news.qq.com/china_index.shtml</link>
        <copyright>Copyright 1998 - 2005 TENCENT Inc. All Rights Reserved</copyright>
        <language>zh-cn</language>
        <generator>www.qq.com</generator>
        <item>
          <title>国资委就“国有企业改革发展”相关问题答记者问</title>
          <link>http://news.qq.com/a/20180310/018317.htm</link>
          <author>www.qq.com</author>
          <category/>
          <pubDate>2018-03-10 14:14:11</pubDate>
          <comments/>
          <description>十三届全国人大一次会议新闻中心将于3月10日15时在梅地亚中心多功能厅举行记者会，邀请国务院国资委主任肖亚庆，副秘书长、新闻发言人彭华岗就“国有企业改革发展”相关问题回答中外记者提问。</description>
        </item>
        <item>
          <title>外媒看两会：中国大胆迈向金融开放 货币政策保持定力</title>
          <link>http://news.qq.com/a/20180310/018312.htm</link>
          <author>www.qq.com</author>
          <category/>
          <pubDate>2018-03-10 14:14:00</pubDate>
          <comments/>
          <description>　　中新网3月10日电全国“两会”继续进行，3月9日上午，十三届全国人大一次会议新闻中心举行关于金融领域热点问题的记者会，外媒对此高度关注。外媒认为，中国进入稳杠杆和逐步调降杠杆阶段，金融监管体制改革呼之欲出；中国对外更“大胆地”开放金融市场的表态，也引起多家外媒关注。　　十三届全国人大一次会议新闻中心</description>
        </item>
        <item>
          <title>田春艳代表：反思莫焕晶案，建全国联网家政服务人员信用平台</title>
          <link>http://news.qq.com/a/20180310/016960.htm</link>
          <author>www.qq.com</author>
          <category/>
          <pubDate>2018-03-10 12:56:34</pubDate>
          <comments/>
          <description>全国两会期间，来自北京代表团的全国人大代表田春艳向十三届全国人大一次会议提交了《关于建立全国联网的家政服务人员信用平台的建议》（下称“《建议》”）。建议中称，622杭州蓝色钱江小区保姆纵火案于2018年2月9日一审公开宣判，被告人莫焕晶被判死刑。此案出来后，社会上除了对受害者哀悼、惋惜、遗憾，对凶手的愤怒、</description>
        </item>
      </channel>
    </rss>
    `
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

func Test_notUTF8(t *testing.T) {
	r := rss.New()

	d := xml.NewDecoder(bytes.NewReader([]byte(notUTF8rss)))
	d.CharsetReader = func(s string, reader io.Reader) (io.Reader, error) {
		return charset.NewReader(reader, s)
	}
	err := d.Decode(r)

	if err != nil {
		t.Fatal(err)
	}

	for _, item := range r.ItemList {
		fmt.Printf("* %s\n%s\n", item.Title, item.Link)
	}
}
