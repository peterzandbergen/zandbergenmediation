package main

import (
	"net/http"
)

const (
	robotText = ``

	// Sitemap xml
	sitemapText = `<?xml version="1.0" encoding="UTF-8"?>
<urlset
      xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
            http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
<url>
  <loc>http://www.zandbergenmediation.nl/index.html</loc>
</url>
<url>
  <loc>http://www.zandbergenmediation.nl/mediation.html</loc>
</url>
<url>
  <loc>http://www.zandbergenmediation.nl/uwmediator.html</loc>
</url>
<url>
  <loc>http://www.zandbergenmediation.nl/nmi.html</loc>
</url>
<url>
  <loc>http://www.zandbergenmediation.nl/contact.html</loc>
</url>
<url>
  <loc>http://www.zandbergenmediation.nl/static/Artikel_Libelle_Vanjefamiliemoetjehethebben.pdf</loc>
</url>
</urlset>


`
)

func handleSitemap(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/xml")
	w.Write([]byte(sitemapText))
}
