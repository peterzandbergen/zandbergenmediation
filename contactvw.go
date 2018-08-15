package main

const contactView = `<!-- Put IE into quirks mode -->
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="description" content="Neem voor vragen en opmerkingen over mediation contact op via email of telefoon">
<title>ZandbergenMediation :: Contact</title>
<link rel="stylesheet" type="text/css" href="css/zandbergenmediation.css" />
<!-- Google Analytics -->
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-45009375-1', 'zandbergenmediation.nl');
  ga('send', 'pageview');
</script>
</head>

<body>
<div id="header"><img alt="Header-left" src="img/header_left.jpg" /><img alt="Logo" src="img/header_logo.jpg" /><img alt="Header-right" src="img/header_right.jpg" /></div>
<div id="footer"><img alt="Footer-left" src="img/footer_left.jpg" /><span class="empty"><br /></span><span class="copy">&copy; ZandbergenMediation 2006</span></div>

<div id="content">
<h1><span class="empty"><br /></span>Contact</h1>
<p>Heeft u vragen en/of opmerkingen? Neem dan contact op. Ik hoor graag van u!</p>
<p>T: +31 (0)6 53 32 66 34<br />
E: <a class="content" title="Neem contact op met ZandbergenMediation" href="mailto: &#103;&#097;&#098;&#121;&#064;&#122;&#097;&#110;&#100;&#098;&#101;&#114;&#103;&#101;&#110;&#109;&#101;&#100;&#105;&#097;&#116;&#105;&#111;&#110;&#046;&#110;&#108;">&#103;&#097;&#098;&#121;&#064;&#122;&#097;&#110;&#100;&#098;&#101;&#114;&#103;&#101;&#110;&#109;&#101;&#100;&#105;&#097;&#116;&#105;&#111;&#110;&#046;&#110;&#108;</a></p>
</p>


</div>
<div id="left">
<div id="menu">
      <h1>Menu</h1>
      <a title="Naar de startpagina" href="index.html">Home</a>
        <a title="Wat is mediation?" href="mediation.html">Mediation</a>
        <a title="Uw mediator Gaby Zandbergen-Vrijens" href="uwmediator.html">Uw mediator</a>
        <a title="Informatie over het Nederlands Mediation Instituut" href="nmi.html">NMI info</a>
        <a class="contactactive" title="Neem contact op met ZandbergenMediation" href="contact.html"><img src="img/arrows.gif" /> Contact</a> 
    </div>
	<div id="searchblock">
      <h1>Zoeken</h1>
		<form action="search.php" method="get" name="searchForm" id="searchForm">
	        <input type="image" name="goBtn" id="goBtn" src="img/go_btn.jpg" alt="Go" /><input type="text" name="query" id="search" value="" action="include/js_suggest/suggest.php" columns="2" autocomplete="off" delay="1500">	
			<input type="hidden" name="search" value="1">
			<a class="singleArrowLink" href="search.php?adv=1">uitgebreid zoeken</a>
        </form>
    </div>
		
	<div style="text-align: center; width: 160px"><p>Libelle - februari 2008:</p>
		<a href="static/Artikel_Libelle_Vanjefamiliemoetjehethebben.pdf" title="Van je familie moet je het hebben - De vier valkuilen van het familiebedrijf" target="_blank"><img src="img/vanjefamiliemoetjehethebben.jpg" /></a>
		<p>m.m.v. uw mediator</p>
    </div>

</div>
</body>
</html>
`
