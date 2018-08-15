package main

const nmiView = `<!-- Put IE into quirks mode -->
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
	"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="description" content="Informatie over NMI, Nederlands Mediation Instituut">
<title>ZandbergenMediation :: NMI info</title>
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
<div id="header"><img alt="Header-left" src="img/header_left.jpg" /><img alt="Logo ZandbergenMediation" src="img/header_logo.jpg" /><img alt="Header-right" src="img/header_right.jpg" /></div>
<div id="footer"><img alt="Footer-left" src="img/footer_left.jpg" /><span class="empty"><br /></span><span class="copy">&copy; ZandbergenMediation 2006</span></div>

<div id="content">
<h1><span class="empty"><br /></span>NMI info</h1>

    <p>ZandbergenMediation is NMI geregistreerd. Het NMI, het Nederlandse Mediation Instituut, is een onafhankelijke stichting met als hoofdtaak de kwaliteitsborging van mediation en de aangesloten en geregistreerde mediators. Verder verstrekt het NMI onafhankelijke informatie en voorlichting over mediation.</p>

      <p><a class="content" target="_blank" title="Bekijk de gedragsregels voor mediators" href="http://www.nmi-mediation.nl/over_mediation/nmi_reglementen_en_modellen/nmi_gedragsregels.php">Gedragsregels mediator <img src="img/arrows.gif" alt="Pijltjes" /></a></p>

      <h3>Klachten</h3>

    <p>Het NMI voorziet in een mogelijkheid klachten te behandelen, middels het klachtenreglement. Ook is voorzien in tuchtrechtspraak.</p>
      <p><a class="content" target="_blank" title="Bekijk het NMI-klachtenreglement" href="http://www.nmi-mediation.nl/over_mediation/niet_tevreden_over_uw_mediator.php">Klachtenreglement <img src="img/arrows.gif" alt="Pijltjes" /></a></p>
 
      <h3>Website NMI</h3>

    <p>Het klachtenreglement, als ook het reglement Stichting Tuchtrechtspraak en verder diverse algemene informatie over mediation is ook rechtstreeks toegankelijk via de website van het NMI:</p>
<p><a class="content" target="_blank" title="Naar de website van het NMI" href="http://www.nmi-mediation.nl">www.nmi-medation.nl <img src="img/arrows.gif" alt="Pijltjes" /></a></p>
    <p>Via deze website kan ook het openbare NMI Register van Mediators worden geraadpleegd.</p>


</div>
<div id="left">
<div id="menu">
      <h1>Menu</h1>
      <a title="Naar de startpagina" href="index.html">Home</a>
        <a title="Wat is mediation?" href="mediation.html">Mediation</a>
        <a title="Uw mediator Gaby Zandbergen-Vrijens" href="uwmediator.html">Uw mediator</a>
        <a class="active" title="Informatie over het Nederlands Mediation Instituut" href="nmi.html"><img src="img/arrows.gif" /> NMI info</a>
        <a class="contact" title="Neem contact op met ZandbergenMediation" href="contact.html">Contact</a> 
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
