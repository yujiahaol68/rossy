package atom_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/yujiahaol68/rossy/atom"
)

const (
	atomContent = `<?xml version="1.0" encoding="utf-8"?>
	<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en-us">
	<title type="text">Blog@Case</title>
	<id>http://blog.case.edu/</id>
	<link rel="alternate" type="application/xhtml+xml" href="http://blog.case.edu/" />
	<link rel="self" type="application/atom+xml" href="http://blog.case.edu/news/feed.atom"/>
	<author>
	<name>jms18</name>
	<uri>http://blog.case.edu/jms18/</uri>
	<email>jeremy.smith@case.edu</email>
	</author>
	<generator uri="http://www.sixapart.com/movabletype/" version="3.121">Movable Type</generator>
	<updated>2006-07-21T17:40:40Z</updated>
	<entry>
	<title type="html">Trackback is Back</title>
	<summary type="text" xml:lang="en">You may have noticed that the trackback functionality of the blog system has been disabled for some time now. We have been engaged in an epic struggle with spammers, and while most of the spammy trackbacks were caught by the...</summary>
	<content type="html">&lt;p&gt;You may have noticed that the &lt;a title=&quot;CaseBlog/FAQ/Tech - CaseWiki&quot; href=&quot;http://wiki.case.edu/CaseBlog/FAQ/Tech#What_is_Trackback&quot;&gt;trackback&lt;/a&gt; functionality of the blog system has been disabled for some time now.  We have been engaged in an epic struggle with spammers, and while most of the &lt;em&gt;spammy&lt;/em&gt; trackbacks were caught by the system, the sheer number of machines attempting to spam us caused an effective &lt;strong&gt;Denial of Service&lt;/strong&gt;, and we were forced to disable trackbacks to keep the blog system operational.&lt;/p&gt;
	
	&lt;p&gt;Well, we&apos;re happy to announce we&apos;ve turned the tide on the spammers and are now blocking those computers that attempt to repeatedly hit the system with trackback or comment spam.  This should also reduce the amount of time you need to go in and spend effort marking icky comments and trackback as spam.&lt;/p&gt;
	
	&lt;p&gt;The new system has been running in a logging mode to gather metrics on what would be acceptable use versus spammer characteristics.  The heuristics used are fairly draconian.  A computer is permanently banned from reaching the blog system if they:&lt;ul&gt;&lt;li&gt;Attempt to submit 5 comments in any 90 minute period and all of those comments end up labelled &quot;moderated&quot; by the other spam measures.&lt;/li&gt;&lt;li&gt;Attempt to submit 2 comments in any 4 hour period from an IP address whose earlier comments were marked as spam by a user of the blog system.&lt;/li&gt;&lt;li&gt;Attempt to trackback to an entry on the blog system 3 times in 90 minute period.&lt;/li&gt;&lt;li&gt;Attempt any trackback from an IP address whose earlier trackback(s) were marked as spam by a user of the blog system.&lt;/li&gt;&lt;/ul&gt;&lt;/p&gt;
	
	&lt;p&gt;If a computer does get blocked, they are given a message telling them so and asking them to email &lt;a href=&quot;mailto:blog-admin@case.edu&quot;&gt;blog-admin@case.edu&lt;/a&gt; to get their computer unblocked.&lt;/p&gt;
	
	&lt;p&gt;As of right now, the new banning system has been running for under 24 hours and over 1000 IPs have been banned.&lt;/p&gt;
	
	&lt;p&gt;We&apos;re winning the war!  Go off and trackback to your heart&apos;s content!&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2006/07#009972</id>
	<link rel="alternate" href="http://blog.case.edu/news/2006/07#009972" type="application/xhtml+xml" hreflang="en" />
	<published>2006-07-21T17:37:19Z</published>
	<updated>2006-07-21T17:40:40Z</updated>
	</entry>
	<entry>
	<title type="html">Yet Another Distributed Attack by a Spammer</title>
	<summary type="text" xml:lang="en">We had another spammer target the system; this time targeting comments. We&apos;ve blocked the spammer and cleaned his comments from the database. If you receive email notifications of new comments on your weblog, you may have received a lot of...</summary>
	<content type="html">&lt;p&gt;We had another spammer target the system; this time targeting comments.  We&apos;ve blocked the spammer and cleaned his comments from the database.  If you receive email notifications of new comments on your weblog, you may have received &lt;strong&gt;a lot&lt;/strong&gt; of emails over last night and this morning.  The offending comments have been deleted.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2006/04#008079</id>
	<link rel="alternate" href="http://blog.case.edu/news/2006/04#008079" type="application/xhtml+xml" hreflang="en" />
	<published>2006-04-08T19:55:07Z</published>
	<updated>2006-04-08T19:54:16Z</updated>
	</entry>
	<entry>
	<title type="html">Trackback Temporarily Disabled</title>
	<summary type="text" xml:lang="en">Trackback is temporarily disabled. A spammer has targetted us with a distributed attack. (We&apos;re currently having a little fun with him at his expense.) Trackback should be turned back on shortly....</summary>
	<content type="html">&lt;p&gt;&lt;a title=&quot;CaseBlog/FAQ/Tech - Edit this page - CaseWiki&quot; href=&quot;http://wiki.case.edu/CaseBlog/FAQ/Tech#What_is_Trackback&quot;&gt;Trackback&lt;/a&gt; is temporarily disabled.  A spammer has targetted us with a distributed attack.  (We&apos;re currently having a little fun with him at his expense.)  Trackback should be turned back on shortly.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2006/04#008038</id>
	<link rel="alternate" href="http://blog.case.edu/news/2006/04#008038" type="application/xhtml+xml" hreflang="en" />
	<published>2006-04-07T04:31:47Z</published>
	<updated>2006-04-07T04:34:07Z</updated>
	</entry>
	<entry>
	<title type="html">We&apos;re Back!</title>
	<summary type="text" xml:lang="en">And a little wary. Things might be a little rough around the edges as we finish getting the new software environment up to snuff....</summary>
	<content type="html">&lt;p&gt;And a little wary.  Things might be a little rough around the edges as we finish getting the new software environment up to snuff.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/11#004845</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/11#004845" type="application/xhtml+xml" hreflang="en" />
	<published>2005-11-18T04:20:53Z</published>
	<updated>2005-11-18T04:21:46Z</updated>
	</entry>
	<entry>
	<title type="html">Intermittent Failures</title>
	<summary type="text" xml:lang="en">The server running the Case Blogging System (and it&apos;s sister site, the Case Wiki) is experiencing intermittent failures. The symptoms of the failures reveal themselves most prominently as the front page and Planet Case not refreshing and per-blog statistics not...</summary>
	<content type="html">&lt;p&gt;The server running the &lt;a title=&quot;Blog@Case&quot; href=&quot;http://blog.case.edu/&quot;&gt;Case Blogging System&lt;/a&gt; (and it&apos;s sister site, the &lt;a title=&quot;Main Page - CaseWiki&quot; href=&quot;http://wiki.case.edu/Main_Page&quot;&gt;Case Wiki&lt;/a&gt;) is experiencing intermittent failures.  The symptoms of the failures reveal themselves most prominently as the front page and &lt;a title=&quot;Planet Case&quot; href=&quot;http://planet.case.edu/&quot;&gt;Planet Case&lt;/a&gt; not refreshing and per-blog statistics not getting updated.  (Other normal means of accessing the blog and the wiki services remain available, though.)&lt;/p&gt;
	
	&lt;p&gt;We are in the process of troubleshooting the root cause of the server&apos;s problem(s), but it seems that we may need to bring the server up and down repeatedly to find what exactly is going on.  In an effort to mitigate down time on the services, we are exploring our options to moving to a new server whilst we troubleshoot the current one.  Please expect some choppy waters as we weigh our options.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/11#004803</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/11#004803" type="application/xhtml+xml" hreflang="en" />
	<published>2005-11-15T19:31:59Z</published>
	<updated>2005-11-15T19:32:49Z</updated>
	</entry>
	<entry>
	<title type="html">Requesting Feedback</title>
	<summary type="text" xml:lang="en">We&apos;ve created an area on the Case Wiki to begin collecting ideas, requests, what-have-you on improvements to be made to the Case Blog system. Specifically, areas of the system that need improvements to help with the ease of use of...</summary>
	<content type="html">&lt;p&gt;We&apos;ve created an area on the &lt;a title=&quot;Main Page - CaseWiki&quot; href=&quot;http://wiki.case.edu/&quot;&gt;Case Wiki&lt;/a&gt; to begin collecting ideas, requests, what-have-you on improvements to be made to the &lt;a title=&quot;Blog@Case&quot; href=&quot;http://blog.case.edu/&quot;&gt;Case Blog system&lt;/a&gt;.  Specifically, areas of the system that need improvements to help with the &lt;a title=&quot;Management Professor Notes II: using the Blog@Case platform&quot; href=&quot;http://blog.case.edu/kep2/2005/10/10/using_the_blogcase_platform&quot;&gt;ease of use&lt;/a&gt; of the system.  So, if you have any ideas of feedback, head on over to &lt;a title=&quot;CaseBlog:Suggestions - CaseWiki&quot; href=&quot;http://wiki.case.edu/CaseBlog:Suggestions&quot;&gt;http://wiki.case.edu/CaseBlog:Suggestions&lt;/a&gt; and leave a note.  Alternatively, you can always &lt;a href=&quot;mailto:blog-admin@case.edu?subject=The%20Case%20Blog%20System%20Is%20Nice%20but%20Could%20be%20Improved%20in%20One%20of%20the%20Following%20Ways&quot;&gt;email us&lt;/a&gt;.  We welcome all kinds of feedback.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/10#003924</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/10#003924" type="application/xhtml+xml" hreflang="en" />
	<published>2005-10-10T17:07:59Z</published>
	<updated>2005-10-10T17:08:25Z</updated>
	</entry>
	<entry>
	<title type="html">Weblogs Being Deleted</title>
	<summary type="text" xml:lang="en">There is an error in the blog system that is causing weblogs to be deleted. As of rate now, we have had three user reports of this. Engineers are tracking down the problem, and we hope to have it fixed...</summary>
	<content type="html">&lt;p&gt;There is an error in the blog system that is causing weblogs to be deleted.  As of rate now, we have had three user reports of this.  Engineers are tracking down the problem, and we hope to have it fixed shortly.&lt;/p&gt;
	
	&lt;p&gt;If this happens to you, please email the &lt;a href=&quot;mailto:blog-admin@case.edu&quot;&gt;Case Blog Administrators&lt;/a&gt;, and we will attempt to restore your blog&apos;s content from backup.&lt;/p&gt;
	
	&lt;p&gt;Thank you, and we are sorry for this inconvenience.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/10#003576</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/10#003576" type="application/xhtml+xml" hreflang="en" />
	<published>2005-10-04T04:23:31Z</published>
	<updated>2005-10-04T04:22:56Z</updated>
	</entry>
	<entry>
	<title type="html">Alumni Access to the Case Blog System</title>
	<summary type="text" xml:lang="en">Alumni have been granted access to the Case Blogging system! As of right now, this is on an experimental basis; so we can determine what load increases this incurs on the server, man-hours, and support. Disregarding all of that disclaimer,...</summary>
	<content type="html">&lt;p&gt;Alumni have been granted access to the Case Blogging system!  As of right now, this is on an experimental basis; so we can determine what load increases this incurs on the server, man-hours, and support.&lt;/p&gt;
	
	&lt;p&gt;Disregarding all of that disclaimer, though... &lt;strong&gt;Welcome Alumni!&lt;/strong&gt;&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002404</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002404" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-21T19:31:59Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Trackback is Back On... This Time, We Mean It</title>
	<summary type="text" xml:lang="en">Okay, [[Trackback]] is back on; and this time we&apos;re serious. We have (for the time being) conquered the spammers (die, spammers, die). If you notice a inordinate amount of [[Trackback]] spam on your blog, however, quickly email the administrators; so...</summary>
	<content type="html">&lt;p&gt;Okay, &lt;a href=&quot;http://wiki.case.edu/Trackback&quot;&gt;Trackback&lt;/a&gt; is back on; and this time we&apos;re &lt;strong&gt;serious&lt;/strong&gt;.  We have (for the time being) conquered the spammers (&lt;em&gt;die, spammers, die&lt;/em&gt;).&lt;/p&gt;
	
	&lt;p&gt;If you notice a inordinate amount of &lt;a href=&quot;http://wiki.case.edu/Trackback&quot;&gt;Trackback&lt;/a&gt; spam on your blog, however, quickly email &lt;a href=&quot;mailto:blog-admin@case.edu&quot;&gt;the administrators&lt;/a&gt;; so we can take care of the problem.&lt;/p&gt;
	
	&lt;p&gt;Thank you, all, for being patient with us during this problem.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002299</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002299" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-14T05:47:47Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Trackbacks (Once Again) Turned Off</title>
	<summary type="text" xml:lang="en">For those keeping score at home, [[Trackback]] has, once again, been turned off for the time being....</summary>
	<content type="html">&lt;p&gt;For those keeping score at home, &lt;a href=&quot;http://wiki.case.edu/Trackback&quot;&gt;Trackback&lt;/a&gt; has, once again, been turned off for the time being.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002292</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002292" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-13T21:47:11Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Trackback Back On</title>
	<summary type="text" xml:lang="en">[[Trackback]] is back on... at least, temporarily while we monitor the situation. If you suddenly notice 10 or more Trackback spams all at once on your blog, please email blog-admin@case.edu to notify us of the problem. And, don&apos;t forget to...</summary>
	<content type="html">&lt;p&gt;&lt;a href=&quot;http://wiki.case.edu/Trackback&quot;&gt;Trackback&lt;/a&gt; is back on... at least, temporarily while we monitor the situation.  If you suddenly notice 10 or more Trackback spams all at once on your blog, please email &lt;a href=&quot;mailto:blog-admin@case.edu?subject=Argh!%20I%20Am%20Getting%20Spammed&quot;&gt;blog-admin@case.edu&lt;/a&gt; to notify us of the problem.  And, don&apos;t forget to despam your blog lest you end up all spammified.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002281</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002281" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-12T20:31:59Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Trackback Temporarily Disabled</title>
	<summary type="text" xml:lang="en">It seems the front layer defense against [[Trackback]] spam has decided to stop working. (There&apos;s a bug somewhere in the code that interacts with the Storable Perl module.) For the time being, trackbacks are disabled....</summary>
	<content type="html">&lt;p&gt;It seems the front layer defense against &lt;a href=&quot;http://wiki.case.edu/Trackback&quot;&gt;Trackback&lt;/a&gt; spam has decided to stop working.  (There&apos;s a bug somewhere in the code that interacts with the &lt;a title=&quot;search.cpan.org: Storable - persistence for Perl data structures&quot; href=&quot;http://search.cpan.org/~ams/Storable-2.15/Storable.pm&quot;&gt;Storable&lt;/a&gt; Perl module.)  For the time being, trackbacks are disabled.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002255</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002255" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-10T21:31:31Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Problems Accessing the Blog Control Panel</title>
	<summary type="text" xml:lang="en">Over the night, some users may have experienced troubles logging in to the blog system. This was caused by an erroneous group in the directory server. Actually, it was caused by two erroneous groups; each claiming to be the one...</summary>
	<content type="html">&lt;p&gt;Over the night, some users may have experienced troubles logging in to the blog system.  This was caused by an erroneous group in the directory server.  Actually, it was caused by &lt;strong&gt;two&lt;/strong&gt; erroneous groups; each claiming to be the one true group that permits or denies access to the Weblog Control Panel.  We removed the more erroneous group and access returned to a status of unfettered.
	&lt;/p&gt;
	
	&lt;p&gt;We now return you to your normal blogging ways.&lt;/p&gt;
	
	&lt;p style=&quot;fontsize: 11px; color: #ccc&quot;&gt;
	This post was brought to you by the letter &lt;strong&gt;E&lt;/strong&gt; for &lt;strong&gt;erroneous&lt;/strong&gt;.
	&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/07#002233</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/07#002233" type="application/xhtml+xml" hreflang="en" />
	<published>2005-07-08T17:59:47Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Doc Oc</title>
	<summary type="text" xml:lang="en">We received the news of Dr. Ocasio&apos;s passing with heavy hearts. To help memorialize his contributions to all of the lives he has touched, we have setup a blog where anyone that knew him can leave their memories....</summary>
	<content type="html">&lt;p&gt;We received the news of Dr. Ocasio&apos;s passing with heavy hearts.  To help memorialize his contributions to all of the lives he has touched, we have setup a &lt;a title=&quot;In Memory of Dr. Ignacio Ocasio&quot; href=&quot;http://blog.case.edu/dococ/&quot;&gt;blog&lt;/a&gt; where anyone that knew him can &lt;a title=&quot;In Memory of Dr. Ignacio Ocasio: Doc Oc&quot; href=&quot;http://blog.case.edu/dococ/2005/05/14/doc_oc&quot;&gt;leave their memories&lt;/a&gt;.&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/05#001692</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/05#001692" type="application/xhtml+xml" hreflang="en" />
	<published>2005-05-16T17:53:13Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	<entry>
	<title type="html">Call for Volunteers to Pilot Group Blogs</title>
	<summary type="text" xml:lang="en">Would you like your department to have a weblog? What about your organization? Or, even, one of your classes? You are in luck because we are looking for volunteers who would like to create group blogs. A group blog is...</summary>
	<content type="html">&lt;p&gt;Would you like your department to have a weblog?  What about your organization?  Or, even, one of your classes?&lt;/p&gt;
	
	&lt;p&gt;You are in luck because we are looking for volunteers who would like to create group blogs.  A group blog is one where there are multiple contributors to the same weblog.  A group blog can be used for announcements, a news site (like the &lt;a title=&quot;Information Technology Services at Case&quot; href=&quot;http://www.case.edu/its&quot;&gt;ITS site&lt;/a&gt;), or an enhanced discussion board meant to foster participation and communication.  Think &lt;code&gt;http://blog.case.edu/depts/mids&lt;/code&gt; or &lt;code&gt;http://blog.case.edu/orgs/cheese_club&lt;/code&gt; or &lt;code&gt;http://blog.case.edu/courses/eces131&lt;/code&gt;; there are a lot of possibilities.&lt;/p&gt;
	
	&lt;p&gt;&lt;br /&gt;
	If you are interested, or you think you may be interested you just aren&apos;t sure if you&apos;re interested and you would like to talk to someone who could get you interested, send an email to the &lt;a href=&quot;mailto:blog-admin@case.edu?subject=About These Group Blogs...&quot;&gt;Case Blog Administrators&lt;/a&gt;.  We&apos;re happy to help you out in starting down this very interesting path of collaboration and improved channels of communication.&lt;/p&gt;
	
	&lt;p&gt;Happy Blogging!,&lt;br /&gt;
	Case Blog Administrators&lt;/p&gt;</content>
	
	<category term="/blognews" scheme="http://blog.case.edu/news/" label="blog-news" />
	
	
	<id>http://blog.case.edu/news/2005/02#000664</id>
	<link rel="alternate" href="http://blog.case.edu/news/2005/02#000664" type="application/xhtml+xml" hreflang="en" />
	<published>2005-02-11T20:31:11Z</published>
	<updated>2005-09-08T18:43:32Z</updated>
	</entry>
	
	</feed>`
)

func Test_parse(t *testing.T) {
	a := atom.New()
	err := xml.Unmarshal([]byte(atomContent), a)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range a.EntryList {
		fmt.Printf("* %s\n%s\n", entry.Title, entry.Link.Href)
	}
}
