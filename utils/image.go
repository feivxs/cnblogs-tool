package utils

import "strings"

func GetImageType(imagePath string) string {
	lastDotIndex := strings.LastIndex(imagePath, ".")
	suffix := imagePath[lastDotIndex:]
	return typeMap[suffix]
}

var typeMap = make(map[string]string, 197)

func init() {
	typeMap[".323"] = "text/h323"
	typeMap[".asx"] = "video/x-ms-asf"
	typeMap[".acx"] = "application/internet-property-stream"
	typeMap[".ai"] = "application/postscript"
	typeMap[".aif"] = "audio/x-aiff"
	typeMap[".aiff"] = "audio/aiff"
	typeMap[".axs"] = "application/olescript"
	typeMap[".aifc"] = "audio/aiff"
	typeMap[".asr"] = "video/x-ms-asf"
	typeMap[".avi"] = "video/x-msvideo"
	typeMap[".asf"] = "video/x-ms-asf"
	typeMap[".au"] = "audio/basic"
	typeMap[".application"] = "application/x-ms-application"
	typeMap[".bin"] = "application/octet-stream"
	typeMap[".bas"] = "text/plain"
	typeMap[".bcpio"] = "application/x-bcpio"
	typeMap[".bmp"] = "image/bmp"
	typeMap[".cdf"] = "application/x-cdf"
	typeMap[".cat"] = "application/vndms-pkiseccat"
	typeMap[".crt"] = "application/x-x509-ca-cert"
	typeMap[".c"] = "text/plain"
	typeMap[".css"] = "text/css"
	typeMap[".cer"] = "application/x-x509-ca-cert"
	typeMap[".crl"] = "application/pkix-crl"
	typeMap[".cmx"] = "image/x-cmx"
	typeMap[".csh"] = "application/x-csh"
	typeMap[".cod"] = "image/cis-cod"
	typeMap[".cpio"] = "application/x-cpio"
	typeMap[".clp"] = "application/x-msclip"
	typeMap[".crd"] = "application/x-mscardfile"
	typeMap[".deploy"] = "application/octet-stream"
	typeMap[".dll"] = "application/x-msdownload"
	typeMap[".dot"] = "application/msword"
	typeMap[".doc"] = "application/msword"
	typeMap[".dvi"] = "application/x-dvi"
	typeMap[".dir"] = "application/x-director"
	typeMap[".dxr"] = "application/x-director"
	typeMap[".der"] = "application/x-x509-ca-cert"
	typeMap[".dib"] = "image/bmp"
	typeMap[".dcr"] = "application/x-director"
	typeMap[".disco"] = "text/xml"
	typeMap[".exe"] = "application/octet-stream"
	typeMap[".etx"] = "text/x-setext"
	typeMap[".evy"] = "application/envoy"
	typeMap[".eml"] = "message/rfc822"
	typeMap[".eps"] = "application/postscript"
	typeMap[".flr"] = "x-world/x-vrml"
	typeMap[".fif"] = "application/fractals"
	typeMap[".gtar"] = "application/x-gtar"
	typeMap[".gif"] = "image/gif"
	typeMap[".gz"] = "application/x-gzip"
	typeMap[".hta"] = "application/hta"
	typeMap[".htc"] = "text/x-component"
	typeMap[".htt"] = "text/webviewhtml"
	typeMap[".h"] = "text/plain"
	typeMap[".hdf"] = "application/x-hdf"
	typeMap[".hlp"] = "application/winhlp"
	typeMap[".html"] = "text/html"
	typeMap[".htm"] = "text/html"
	typeMap[".hqx"] = "application/mac-binhex40"
	typeMap[".isp"] = "application/x-internet-signup"
	typeMap[".iii"] = "application/x-iphone"
	typeMap[".ief"] = "image/ief"
	typeMap[".ivf"] = "video/x-ivf"
	typeMap[".ins"] = "application/x-internet-signup"
	typeMap[".ico"] = "image/x-icon"
	typeMap[".jpg"] = "image/jpeg"
	typeMap[".jfif"] = "image/pjpeg"
	typeMap[".jpe"] = "image/jpeg"
	typeMap[".jpeg"] = "image/jpeg"
	typeMap[".js"] = "application/x-javascript"
	typeMap[".lsx"] = "video/x-la-asf"
	typeMap[".latex"] = "application/x-latex"
	typeMap[".lsf"] = "video/x-la-asf"
	typeMap[".manifest"] = "application/x-ms-manifest"
	typeMap[".mhtml"] = "message/rfc822"
	typeMap[".mny"] = "application/x-msmoney"
	typeMap[".mht"] = "message/rfc822"
	typeMap[".mid"] = "audio/mid"
	typeMap[".mpv2"] = "video/mpeg"
	typeMap[".man"] = "application/x-troff-man"
	typeMap[".mvb"] = "application/x-msmediaview"
	typeMap[".mpeg"] = "video/mpeg"
	typeMap[".m3u"] = "audio/x-mpegurl"
	typeMap[".mdb"] = "application/x-msaccess"
	typeMap[".mpp"] = "application/vnd.ms-project"
	typeMap[".m1v"] = "video/mpeg"
	typeMap[".mpa"] = "video/mpeg"
	typeMap[".me"] = "application/x-troff-me"
	typeMap[".m13"] = "application/x-msmediaview"
	typeMap[".movie"] = "video/x-sgi-movie"
	typeMap[".m14"] = "application/x-msmediaview"
	typeMap[".mpe"] = "video/mpeg"
	typeMap[".mp2"] = "video/mpeg"
	typeMap[".mov"] = "video/quicktime"
	typeMap[".mp3"] = "audio/mpeg"
	typeMap[".mpg"] = "video/mpeg"
	typeMap[".ms"] = "application/x-troff-ms"
	typeMap[".nc"] = "application/x-netcdf"
	typeMap[".nws"] = "message/rfc822"
	typeMap[".oda"] = "application/oda"
	typeMap[".ods"] = "application/oleobject"
	typeMap[".pmc"] = "application/x-perfmon"
	typeMap[".p7r"] = "application/x-pkcs7-certreqresp"
	typeMap[".p7b"] = "application/x-pkcs7-certificates"
	typeMap[".p7s"] = "application/pkcs7-signature"
	typeMap[".pmw"] = "application/x-perfmon"
	typeMap[".ps"] = "application/postscript"
	typeMap[".p7c"] = "application/pkcs7-mime"
	typeMap[".pbm"] = "image/x-portable-bitmap"
	typeMap[".ppm"] = "image/x-portable-pixmap"
	typeMap[".pub"] = "application/x-mspublisher"
	typeMap[".pnm"] = "image/x-portable-anymap"
	typeMap[".png"] = "image/png"
	typeMap[".pml"] = "application/x-perfmon"
	typeMap[".p10"] = "application/pkcs10"
	typeMap[".pfx"] = "application/x-pkcs12"
	typeMap[".p12"] = "application/x-pkcs12"
	typeMap[".pdf"] = "application/pdf"
	typeMap[".pps"] = "application/vnd.ms-powerpoint"
	typeMap[".p7m"] = "application/pkcs7-mime"
	typeMap[".pko"] = "application/vndms-pkipko"
	typeMap[".ppt"] = "application/vnd.ms-powerpoint"
	typeMap[".pmr"] = "application/x-perfmon"
	typeMap[".pma"] = "application/x-perfmon"
	typeMap[".pot"] = "application/vnd.ms-powerpoint"
	typeMap[".prf"] = "application/pics-rules"
	typeMap[".pgm"] = "image/x-portable-graymap"
	typeMap[".qt"] = "video/quicktime"
	typeMap[".ra"] = "audio/x-pn-realaudio"
	typeMap[".rgb"] = "image/x-rgb"
	typeMap[".ram"] = "audio/x-pn-realaudio"
	typeMap[".rmi"] = "audio/mid"
	typeMap[".ras"] = "image/x-cmu-raster"
	typeMap[".roff"] = "application/x-troff"
	typeMap[".rtf"] = "application/rtf"
	typeMap[".rtx"] = "text/richtext"
	typeMap[".sv4crc"] = "application/x-sv4crc"
	typeMap[".spc"] = "application/x-pkcs7-certificates"
	typeMap[".setreg"] = "application/set-registration-initiation"
	typeMap[".snd"] = "audio/basic"
	typeMap[".stl"] = "application/vndms-pkistl"
	typeMap[".setpay"] = "application/set-payment-initiation"
	typeMap[".stm"] = "text/html"
	typeMap[".shar"] = "application/x-shar"
	typeMap[".sh"] = "application/x-sh"
	typeMap[".sit"] = "application/x-stuffit"
	typeMap[".spl"] = "application/futuresplash"
	typeMap[".sct"] = "text/scriptlet"
	typeMap[".scd"] = "application/x-msschedule"
	typeMap[".sst"] = "application/vndms-pkicertstore"
	typeMap[".src"] = "application/x-wais-source"
	typeMap[".sv4cpio"] = "application/x-sv4cpio"
	typeMap[".tex"] = "application/x-tex"
	typeMap[".tgz"] = "application/x-compressed"
	typeMap[".t"] = "application/x-troff"
	typeMap[".tar"] = "application/x-tar"
	typeMap[".tr"] = "application/x-troff"
	typeMap[".tif"] = "image/tiff"
	typeMap[".txt"] = "text/plain"
	typeMap[".texinfo"] = "application/x-texinfo"
	typeMap[".trm"] = "application/x-msterminal"
	typeMap[".tiff"] = "image/tiff"
	typeMap[".tcl"] = "application/x-tcl"
	typeMap[".texi"] = "application/x-texinfo"
	typeMap[".tsv"] = "text/tab-separated-values"
	typeMap[".ustar"] = "application/x-ustar"
	typeMap[".uls"] = "text/iuls"
	typeMap[".vcf"] = "text/x-vcard"
	typeMap[".wps"] = "application/vnd.ms-works"
	typeMap[".wav"] = "audio/wav"
	typeMap[".wrz"] = "x-world/x-vrml"
	typeMap[".wri"] = "application/x-mswrite"
	typeMap[".wks"] = "application/vnd.ms-works"
	typeMap[".wmf"] = "application/x-msmetafile"
	typeMap[".wcm"] = "application/vnd.ms-works"
	typeMap[".wrl"] = "x-world/x-vrml"
	typeMap[".wdb"] = "application/vnd.ms-works"
	typeMap[".wsdl"] = "text/xml"
	typeMap[".xap"] = "application/x-silverlight-app"
	typeMap[".xml"] = "text/xml"
	typeMap[".xlm"] = "application/vnd.ms-excel"
	typeMap[".xaf"] = "x-world/x-vrml"
	typeMap[".xla"] = "application/vnd.ms-excel"
	typeMap[".xls"] = "application/vnd.ms-excel"
	typeMap[".xof"] = "x-world/x-vrml"
	typeMap[".xlt"] = "application/vnd.ms-excel"
	typeMap[".xlc"] = "application/vnd.ms-excel"
	typeMap[".xsl"] = "text/xml"
	typeMap[".xbm"] = "image/x-xbitmap"
	typeMap[".xlw"] = "application/vnd.ms-excel"
	typeMap[".xpm"] = "image/x-xpixmap"
	typeMap[".xwd"] = "image/x-xwindowdump"
	typeMap[".xsd"] = "text/xml"
	typeMap[".z"] = "application/x-compress"
	typeMap[".zip"] = "application/x-zip-compressed"
	typeMap[".*"] = "application/octet-stream"
}