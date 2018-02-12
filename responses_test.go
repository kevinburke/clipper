package clipper

var dashboard = []byte(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head>
	<title>
	    		My Clipper
	    	</title>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />

	<style type="text/css">
.clearfix:after {
	content: ".";
	display: block;
	height: 0;
	clear: both;
	visibility: hidden;
}
</style>

	<!-- Migrated Jquery 1.4.4 to 1.8.3. Solves Security Issue: Cross-site scripting (XSS) vulnerability
	in jQuery before 1.6.3, when using location.hash to select elements, allows remote attackers to inject
	arbitrary web script or HTML via a crafted tag.	-->

	<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/swfobject.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/CSJSRequestObject.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/switchcontent.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/switchicon.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/tabcontent.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/helpBox.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/validation.js.jsf?ln=js"></script><script type="text/javascript" src="/ClipperCard/javax.faces.resource/checkOut.js.jsf?ln=js"></script>


	<script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/jquery-cookie/1.4.1/jquery.cookie.js"></script>

	<!-- Begin MIG HTML add for Clipper eyebrow --><script type="text/javascript" src="/ClipperCard/javax.faces.resource/clipper-eyebrow.js.jsf?ln=js"></script>


	<style type="text/css">
#header h1 a {
	height: 91px;
	margin: 10px 0 10px 20px;
}

#header {
	height: 110px;
}

ul#topNav {
	top: 81px;
}

#myRollover {
	top: 81px;
}
</style>



	<!-- Migrated Jquery UI 1.8.18 to 1.10.4. Solves Security Issue: Cross-site scripting (XSS)
	 vulnerability in jquery.ui.dialog.js in the Dialog widget in jQuery UI before 1.10.0
	 allows remote attackers to inject arbitrary web script or HTML via the title option.	 -->
	<script src="//www.google.com/recaptcha/api.js"></script>
	<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jqueryui/1.10.4/jquery-ui.min.js"></script>
	<script type="text/javascript" charset="utf-8">
		api_gallery = [ 'test1.php?iframe=true&width=600&height=300' ];
		api_titles = [ '', ];
		api_descriptions = [ '' ];
	</script>
	<script language="JavaScript">
		function setCookie(name, value, expires) {
			document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
			if (expires)
				document.cookie = name + "=" + value + ";expires=" + expires
						+ "; path=/";

		}
		function transitSwitch(divToShow) {
			$('div.transitContent').hide();
			$('#' + divToShow).show();
		}

		function menuChange() {
			this.style.backgroundColor = "blue";
		}

		$(function() {

			//break out of frames, if in a frame
			if (top.location != self.location) {
				top.location = self.location.href
			}

			if (typeof jsf !== 'undefined') {
				jsf.ajax.addOnEvent(function(data) {
					var ajaxstatus = data.status; // Can be "begin", "complete" and "success"
					var ajaxloader = document.getElementById("loadingIcon");

					switch (ajaxstatus) {
					case "begin": // This is called right before ajax request is been sent.
						ajaxloader.style.display = 'block';
						break;

					case "complete": // This is called right after ajax response is received.
						ajaxloader.style.display = 'none';
						break;

					case "success": // This is called when ajax response is successfully processed.
						// NOOP.
						break;
					}
				});
			}
			transitSwitch('transitContent1');
		});
	</script><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/prettyPhoto.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/thickbox.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/print2011.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/redesign2010.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/components.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/components2011.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/screen.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/screen2011.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/jquery-ui-1.8.18.custom.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/eyebrow.css.jsf?ln=css" /><script type="text/javascript" src="/ClipperCard/javax.faces.resource/jsf.js.jsf?ln=javax.faces"></script></head>

<body id="myClipperPage" class="myClipper orderCard">
<form id="mainForm" name="mainForm" method="post" action="/ClipperCard/dashboard.jsf" enctype="application/x-www-form-urlencoded">
<input type="hidden" name="mainForm" value="mainForm" />

		<div id="container">
			<div id="wrapper">
				<div id="header">
	<script type="text/javascript">
		$(function(){
			$("#btnSearch").click(function(){
				window.location.href = "/ClipperWeb/search.do?searchString=" + $("#searchText").val();
			});
		})
	</script>
	<div class="header"></div>


	<div id="header">
		<div id="skipto">
			<a href="#mainContent" class="skip_nav" accesskey="z">Skip to
				Content</a>
		</div>
		<div id="headerSearch">
			<label id="searchString" for="searchText">Search</label> <input id="searchText" /> <input value="Go" id="btnSearch" class="buttonSearch" style="padding-left: 12px;" /><br />
			<div style="color: red; font-size: .75em; text-align: right; padding-right: 30px;">
			</div>
		</div>
			<div id="logOffBox">
				Hello,
				KEVIN
				<a href="../ClipperCard/logout.jsf">Log Out</a>
			</div>
		<h1>
			<a href="/ClipperWeb/index.do">Clipper</a>
		</h1>
		<ul id="topNav">
			<li id="home"><a href="../ClipperWeb/index.do" accesskey="2" title="Home">Home</a></li>
			<li id="about"><a href="../ClipperWeb/whatsTranslink.do" accesskey="2" title="About Clipper page">About Clipper</a></li>
			<li id="use"><a href="../ClipperWeb/useTranslink.do" accesskey="3" title="How to use Clipper page">Use Clipper</a></li>
			<li id="get"><a href="../ClipperWeb/getTranslink.do" accesskey="4" title="How to get Clipper page">Get Clipper</a></li>
		</ul>
		<script type="text/javascript">

		$(document).ready(function() {

		 // toggles the slickbox on clicking the noted link
		  $('#myclipperModuleRollover').click(function(){
			  	$("#myBox").slideToggle(150);
			  	$("#myBox input.login").val("");
			})

			$("#myclipperLoggedIn").click(function(){
				$("#myLoggedInBox").slideToggle(150);
			});

			//show the correct login/myClipper div
			if("true" == "true"){
				$("#myRollover1").hide();
				$("#myRollover2").show();
			}
			else{
				if("false" == "true"){
					$("#myRollover1").hide();
					$("#myRollover2").show();
				}
				else{
					$("#myRollover1").show();
					$("#myRollover2").hide();
				}
			}
			if(false)
				$("#myBox").show();

			$("input.login").live("keydown",function(e){
				if(e.keyCode==13){
					e.preventDefault();
					$("#mainForm\\:submitLogin").click();
				}
			});

		});

		</script><div id="mainForm:loginView">

			<div id="myRollover1">

				<h1 id="myclipperModuleRollover">Toggle the box</h1>

				<div id="myBox" style="display: none;">
					<div class="myclipperModuleBody">
						<div style="text-align: left; font-weight: bold; padding-top: 5px;">
							<p>Secure Login</p>
						</div>
						<div style="text-align: left;">
							<label>Email Address:</label>
						</div><input id="mainForm:username" type="text" name="mainForm:username" value="kev@inburke.com" class="login" />
						<div style="text-align: left;">
							<label>Password:</label>
						</div><input id="mainForm:password" type="password" name="mainForm:password" value="" maxlength="30" class="login" />
						<p class="formSmall">
							<a href="forgottenPassword.jsf" style="color: white;" title="Forgot your password?">Forgot your password?</a>
						</p><span id="mainForm:renderError"></span><input id="mainForm:submitLogin" type="submit" name="mainForm:submitLogin" value="Login" class="buttonLogin" onclick="mojarra.ab(this,event,'action','mainForm:username mainForm:password','mainForm:renderError');return false" />
					</div>
					<!-- END MY CLIPPER MODULE -->
					<div class="myclipperModuleBodyBottom"></div>
				</div>
			</div>
			<div id="myRollover2" style="display: none">

				<h1 id="myclipperLoggedIn">Toggle the box</h1>

				<div id="myLoggedInBox" style="display: none;">
					<div class="myclipperModuleBody">
						<ul>
							<li><a href="dashboard.jsf">Manage account</a></li>
							<li><a href="dashboard.jsf">Check my balance</a></li>
							<li><a id="mainForm:addValueLinkOne" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:addValueLinkOne':'mainForm:addValueLinkOne'},'');return false">Add value</a></li>
							<li><a id="mainForm:replaceCardLink" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:replaceCardLink':'mainForm:replaceCardLink'},'');return false">Replace my card</a></li>
							<li><a href="/ClipperWeb/damagedCard.do">Support</a></li>
							<li><a href="logout.jsf">Log Out</a></li>
						</ul>
					</div>
					<!-- END MY CLIPPER MODULE -->
					<div class="myclipperModuleBodyBottom"></div>
				</div>


			</div></div>
	</div>

	<!-- /end #header -->
				</div>
				<div id="leftNav">

	<script type="text/javascript">
	$(function(){
		//get the request url, find the a with the href = request url and apply style navSelected to parent li
		//remove navSelected from any other li values that have it
		var url = window.location.href;
		var endIdx = url.indexOf("?",0);
		var page = url.substring(url.lastIndexOf("/")+1,endIdx==-1?url.length:endIdx);
		endIdx = document.referrer.indexOf("?",0);
		var source =  document.referrer.substring(document.referrer.lastIndexOf("/")+1,endIdx==-1?document.referrer.length:endIdx);


		// added this to show left nav Auotload information box - Indira
		var isShowAutoloadBox = false;
		if(page=="order.jsf" || (page =="manage.jsf" && isShowAutoloadBox )){
			$("#leftnavAutoloadInfoBox").show();
		}else{
			$("#leftnavAutoloadInfoBox").hide();
		}


		//Exception cases
		if(page == "cardHistory.jsf")
			page = "dashboard.jsf";
		else if(page == "checkout.jsf" && source == "order.jsf")
			page = "order.jsf";
		else if(page == "checkout.jsf" && source == "manage.jsf")
			page = "manage.jsf"
		else if(page == "passwordChangeSuccessful.jsf")
			page = "changePassword.jsf"
		else if(page == "lostStoleConfirmation.jsf")
			page = "lostCard.jsf";

		//end exception cases

		$("#leftNavModule li").removeClass("navSelected");
		$("#leftNavModule li[name="+"'" + page+"'" +"]").addClass("navSelected").parent().addClass("navSelected");

		if("true" == "true")
			$("#logOff").show();
		else
			$("#logOff").hide();
	});
	</script>

	<div id="leftNav">
		<ul id="leftNavModule">
					<li id="myAccount" class="navSelected" name="dashboard.jsf"><a href="dashboard.jsf" title="My Account">My account</a></li>
					<li id="orderCard" name="order.jsf"><a id="mainForm:orderNewCardLink" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:orderNewCardLink':'mainForm:orderNewCardLink'},'');return false">Order new card</a></li>
					<li id="addValue" name="manage.jsf"><a id="mainForm:addValueLink" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:addValueLink':'mainForm:addValueLink'},'');return false">Add value to card</a></li>
					<li id="registerCard" name="register.jsf"><a id="mainForm:registerCardLink" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:registerCardLink':'mainForm:registerCardLink'},'');return false">Register a card</a></li>

					<li id="editProfile" name="editProfile.jsf"><a href="editProfile.jsf" title="Edit profile Information">Profile information</a></li>
							<li id="editPayment" name="editPayment.jsf"><a id="mainForm:editFundingSourceLineTwo" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:editFundingSourceLineTwo':'mainForm:editFundingSourceLineTwo'},'');return false">Payment information</a></li>
					<li id="lostCard" name="lostCard.jsf"><a href="lostCard.jsf" title="Report lost, stolen or damaged card">Report lost, stolen or damaged card</a></li>
					<li id="unregisterCardOne" name="unregister.jsf"><a id="mainForm:unregisterCardTwo" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:unregisterCardTwo':'mainForm:unregisterCardTwo'},'');return false">Deregister card</a></li>
					<li id="changePassword" name="changePassword.jsf"><a href="changePassword.jsf" title="Change Password">Change password</a></li>
					<li id="customerSupport" name="customerSupport.jsf"><a href="../ClipperWeb/damagedCard.do" title="Customer Support">Customer support</a></li>
					<li id="logOff" style="display:none;"><a id="mainForm:logOutOne" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:logOutOne':'mainForm:logOutOne'},'');return false">Log out</a></li>

		</ul>
		<ul id="leftNavModuleEnd">
			<li> </li>
		</ul>

		<div id="leftnavAutoloadInfoBox">

		</div>
	</div>
				</div>
				<div id="mainContent"><input type="hidden" name="mainForm:j_idt57" value="0" />
			<script type="text/javascript">

					var curDisableLnk = null;
 					var curSerial = null;
					$(function(){

						 $(".closeWindow").click(function() {

				                $("#dialogOptIn").dialog("close");

				                sessionStorage.setItem("runOnceOptin", true);
				            });

						 $(".closeWindowAndDS").click(function() {
			                $("#dialogOptIn").dialog("close");
			                sessionStorage.setItem("runOnceOptin", true);
			                $("#mainForm\\:doNotShowOptIn").click();

			            });

						$("div.toggleCardInfo").live("click",function(){
							if($(this).html() == "-"){
								$(this).parent().next().hide();
								$(this).html("+")
							}
							else{
								$(this).parent().next().show();
								$(this).html("-")
							}
							setExpandCollapseVis();
						});

						$(".changeNameLink").click(function(){
							var name = $(this).parent().prev().children().first().html();
							$(this).parent().prev().find('.editName').children().first().val(name);
							$(this).parent().parent().find(".displayName").hide();
							$(this).parent().parent().find(".editName").show();
							$(this).parent().parent().find(".editName").children().eq(0).focus();
						});

						$(".cancelEdit").click(function(){
							$(this).parent().parent().parent().find(".displayName").show();
							$(this).parent().parent().parent().find(".editName").hide();
						});

						$(".saveEdit").click(function(){
							var name = $(this).parent().parent().prev().prev().find('.editName').children().first().val();
							$(this).parent().parent().prev().prev().children().first().html(name);
							$(this).parent().parent().parent().find(".displayName").show();
							$(this).parent().parent().parent().find(".editName").hide();
						});

						$("#dialogGetSecurityQuestion").dialog({
							autoOpen:false,
							resizable:false,
							modal:true,
							width:450,
							buttons:{
								"Save":function(){
									$("input.securityAnswer").focus();
									$("input.securityAnswer").blur();

									var errs = $("#dialogGetSecurityQuestion div.boxX").length;
									if(errs == 0){
										$("#securityData input.securityAnswer").val($("#txtSecAnswer").val());
										$("#securityData input.securityQuestion").val($("select.ddlQuestion option:selected").text())
										$(this).dialog("close");
										$("#saveBtn").find("input").click();
									}
								}
							}
						});



						$("#dialogConfirmDisableAutoload").dialog({
							autoOpen:false,
							resizable:false,
							modal:true,
							width:250,
							buttons:{
								"Yes": function()
								{
									$("#" + curDisableLnk).next().click();
									$("#" + curDisableLnk).next().remove();
									$("#" + curDisableLnk).remove();
									$(this).dialog("close")
								},
								"No": function()
								{
									$(this).dialog("close")
								}
							},
							beforeClose: function(){
								curDisableLnk = null;
							}
						});

						$("#dialogOptIn").dialog({
							autoOpen:false,
							resizable:false,
							modal:true,
							dialogClass: 'dashOptDialog',
							width:450,
							beforeClose: function(){
								sessionStorage.setItem("runOnceOptin", true);
							}
						});



						if(false){
							$("#dialogGetSecurityQuestion").dialog("open");
						}

						if(true){
							 if (!sessionStorage.getItem("runOnceOptin")) {
								 $("#dialogOptIn").dialog("open");
							 }
						}


						$("span.disableAutoloadSpan").click(function(){
							curDisableLnk = $(this).prop("id");
							$("#dialogConfirmDisableAutoload").dialog("open");
						});

						$("#dialogNoHistory").dialog({
							autoOpen:false,
							resizable:false,
							modal:true,
							width:500,
							buttons:{
								"Ok": function()
								{
									$(this).dialog("close");
								}
							}
						});

						$("span.changeAutoloadCashAmt").live("click",function(){
							curChangeAutoloadCashLnk = $(this).prop("id");
							$("#newEcashAmt").val("");
							$("#dialogChangeAutoloadAmt").dialog("open");
						})

						$("span.changeAutoloadParkingAmt").live("click",function(){
							curChangeAutoloadParkingLnk = $(this).prop("id");
							$("#newParkingPurseAmt").val("");
							$("#dialogChangeAutoloadParkingPurseAmt").dialog("open");
						})



						$('#dialogChangeAutoloadAmt').keyup(function(e) {
						    if (e.keyCode == 13) {
						    	$("#mainForm\\:newEcashAmtVal").val($("#newEcashAmt").val())
								$("#" + curChangeAutoloadCashLnk).next().click();
								$(this).dialog("close");
						    }
						});

						$('#dialogChangeAutoloadParkingPurseAmt').keyup(function(e) {
						    if (e.keyCode == 13) {
						    	$("#mainForm\\:newParkingPurseAmtVal").val($("#newParkingPurseAmt").val())
								$("#" + curChangeAutoloadParkingLnk).next().click();
								$(this).dialog("close");
						    }
						});

						$("#dialogChangeAutoloadAmt").dialog({
							autoOpen:false,
							width:450,
							modal:true,
							resizable:false,
							buttons:
							{
								"Save":function(){
									$("#mainForm\\:newEcashAmtVal").val($("#newEcashAmt").val())
									$("#" + curChangeAutoloadCashLnk).next().click();
									$(this).dialog("close");
								},
								"Cancel":function(){
									$(this).dialog("close");
								}
							}
						})

						$("#dialogChangeAutoloadParkingPurseAmt").dialog({
							autoOpen:false,
							width:450,
							modal:true,
							resizable:false,
							buttons:
							{
								"Save":function(){
									$("#mainForm\\:newParkingPurseAmtVal").val($("#newParkingPurseAmt").val())
									$("#" + curChangeAutoloadParkingLnk).next().click();
									$(this).dialog("close");

								},
								"Cancel":function(){
									$(this).dialog("close");
								}
							}
						})

					});

					function seeHistory(){
						var nowHrs = new Date().getHours();
						if(nowHrs >= 1  && nowHrs < 3){
							$("#dialogNoHistory").dialog("open");
							return false;
						}
						else
							return true;
					}

					function expandCards(){
						$('div.cardData').show();
						$('div.toggleCardInfo').html('-');
						setExpandCollapseVis();
					}
					function collapseCards(){
						$('div.cardData').hide();
						$('div.toggleCardInfo').html('+');
						setExpandCollapseVis();
					}
					function setExpandCollapseVis(){
						if($("div.cardData:visible").length == $("div.cardData").length){
							$("#collapseLink").show();
							$("#expandLink").hide();
						}
						else{
							$("#collapseLink").hide();
							$("#expandLink").show();
						}
					}

					function setPaymentExpandCollapseVis(){
						if($("div.paymentData:visible").length == $("div.paymentData").length){
							$("#collapsePaymentLink").show();
							$("#expandPaymentLink").hide();
						}
						else{
							$("#collapsePaymentLink").hide();
							$("#expandPaymentLink").show();
						}
					}

					function expandPayment(){
						$('div.paymentData').show();
						setPaymentExpandCollapseVis();
					}
					function collapsePayment(){
						$('div.paymentData').hide();
						setPaymentExpandCollapseVis();
					}

					function setProfileExpandCollapseVis(){
						if($("div.profileData:visible").length == $("div.profileData").length){
							$("#collapseProfileLink").show();
							$("#expandProfileLink").hide();
						}
						else{
							$("#collapseProfileLink").hide();
							$("#expandProfileLink").show();
						}
					}

					function expandProfile(){
						$('div.profileData').show();
						setProfileExpandCollapseVis();
					}
					function collapseProfile(){
						$('div.profileData').hide();
						setProfileExpandCollapseVis();
					}




		    	</script>
			<style type="text/css">
.ui-dialog .ui-dialog-buttonpane button {
	font-size: .9em !important;
}

.ui-dialog .ui-dialog-title {
	float: none !important;
	color: #0069a6 !important;
}

.ui-dialog .dashOptDialog .ui-dialog-titlebar {
	color: #0069a6 !important;
	font-size: inherit;
}
</style>


			<div id="mainContent">
				<div class="prop500"></div>
				<div class="contentSubHeader">
					<p>My Account</p>
				</div>
				<!--MESSAGES-->
				<div class="stepHeader" style="display: none;">
					<div class="title">
						<p>Important Messages</p>
					</div>
				</div><span id="mainForm:pnlErr" style="color:red"></span>
				<!--BOX 1-->
				<div class="greyBox" style="display: none;">
					<p class="errorText">In this box, Clipper would send messages
						about credit cards that are about to expire, failed Autoload
						payments, offer customers the ability to set up a secondary
						payment method if they don't have one set up, etc.</p>
				</div>
				<div class="spacer"></div>
				<!--END MESSAGES-->
				<!-- introduction started -->
				<div id="overviewcontent1">
					<p>Welcome! My Clipper is a secure place to add value to your
						card, set up Autoload, manage your account information, and more.</p>
				</div>
				<!-- introduction finished -->
				<!--PROFILE INFORMATION-->
				<div class="stepHeader">
					<div class="title">
						<p>My Profile Information</p>
					</div>
					<div class="policy">
						<p>
							<span id="collapseProfileLink" style="cursor: pointer; text-decoration: underline; font-size: .75em; color: white;" onclick="collapseProfile();">Collapse </span> <span id="expandProfileLink" style="cursor: pointer; display: none; text-decoration: underline; font-size: .75em; color: white;" onclick="expandProfile();">Expand </span>
						</p>
					</div>
					<!--&lt;div style=&quot;float: right;&quot; class=&quot;helpToggle helpToggleInactive&quot;
						onclick=&quot;showHelp('helpContainerOne', this);&quot;&gt;&lt;/div&gt;-->
				</div>
				<!--BOX 1-->
				<div class="greyBox  profileData">
					<div class="infoDiv">
						<div class="fieldName">Name on Account:</div>
						<div class="fieldData">KEVIN BURKE</div>
					</div>

					<div class="spacer"></div>

					<div class="infoDiv">
						<div class="fieldName">Email:</div>
						<div class="fieldData">kev@inburke.com</div>
					</div>

					<div class="spacer"></div>

					<div class="infoDiv">
						<div class="fieldName">Address:</div>
						<div class="fieldData">639 Old County Road</div>
							<div class="fieldName"></div>
							<div class="fieldData">Apartment 23</div>
					</div>

					<div class="infoDiv">
						<div class="fieldName"></div>
						<div class="fieldData">Belmont,
							CA</div>
					</div>

					<div class="infoDiv">
						<div class="fieldName"></div>
						<div class="fieldData">94002</div>
					</div>

					<div class="spacer"></div>

					<div class="infoDiv">
						<div class="fieldName">Phone:</div>
						<div class="fieldData">925-271-7005</div>
					</div>

					<div class="infoDiv">
						<div class="fieldName">Email Updates:</div>
						<div class="fieldData">No</div>
					</div>

					<div class="spacer"></div>

					<div class="infoDiv">
						<div class="fieldName"></div>
						<div class="fieldData">
							<a href="editProfile.jsf?faces-redirect=true">Edit My Profile
								Information</a>
						</div>
					</div>
				</div>
				<div class="spacer"></div><a id="mainForm:doNotShowOptIn" href="#" style="display:none" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:doNotShowOptIn':'mainForm:doNotShowOptIn'},'');return false"></a>
				<!--END PROFILE INFORMATION--><span id="mainForm:paymentInfoOnFile">
					<!-- &lt;c:if test=&quot;true&quot;&gt; -->
					<!--payment INFORMATION-->
					<div class="stepHeader">
						<div class="title">
							<p>My Payment Information</p>
						</div>
						<div class="policy">
							<p>
								<span id="collapsePaymentLink" style="cursor: pointer; text-decoration: underline; font-size: .75em; color: white;" onclick="collapsePayment();">Collapse </span> <span id="expandPaymentLink" style="cursor: pointer; display: none; text-decoration: underline; font-size: .75em; color: white;" onclick="expandPayment();">Expand </span>
							</p>
						</div>
					</div>
						<span style="align: justify; color: red; font-size: .65em;">
						</span>
					<!--BOX 1-->

					<div class="greyBox paymentData">
						<div class="infoDiv">
							<div class="fieldName">Primary:</div>
							<div class="fieldData">VISA XXXX-XXXX-XXXX-1564</div>
						</div>


						<div class="infoDiv">
							<div class="fieldName"></div>
							<div class="fieldData">Exp.
								01/22

							</div>
							<div class="fieldName"></div><span class="fieldData"><a id="mainForm:editPrimaryFundingOne" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:editPrimaryFundingOne':'mainForm:editPrimaryFundingOne'},'');return false">Edit My Primary Funding</a></span>
						</div>

						<!-- Alternate Payment method is removed for the time being, as per HPQC Ticket 44163 -->
								<div class="spacer"></div>

								<div class="infoDiv">
									<div class="fieldName">Backup:</div>
									<div class="fieldData">NA</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName"></div>
									<div class="fieldData">
									</div>
									<div class="fieldName"></div><span class="fieldData">
										<div style="float: right;" class="helpToggle helpToggleInactive" onclick="showHelp('helpContainerBackup', this);"></div><a id="mainForm:editBackupFundingTwo" href="#" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:editBackupFundingTwo':'mainForm:editBackupFundingTwo'},'');return false">Add My Backup Funding</a></span>



								</div>
								<div class="spacer"></div>
					</div>
					<div class="spacer"></div>
					<!-- &lt;/c:if&gt; --></span>
				<!--END payment INFORMATION-->



				<!--YOUR CLIPPER CARDS-->
				<div class="stepHeader">
					<div class="title">
						<p>My Clipper Cards</p>
					</div>
					<div class="policy">
						<p>
							<span id="collapseLink" style="cursor: pointer; text-decoration: underline; font-size: .75em; color: white;" onclick="collapseCards();">Collapse All</span> <span id="expandLink" style="cursor: pointer; display: none; text-decoration: underline; font-size: .75em; color: white;" onclick="expandCards();">Expand All</span>
						</p>
					</div>
				</div>
				<!--CARD 1-->
				<div class="greyBox2">
						<div class="darkGreyCardHeader">
							<div class="plus toggleCardInfo" style="display: block;">-</div><div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Card Nickname:</div>
									<div class="fieldData field90">
										<span class="displayName">Personal</span><span class="editName" style="display: none"><input id="mainForm:j_idt79:0:cardName" type="text" name="mainForm:j_idt79:0:cardName" value="Personal" /></span>
									</div>

									<div class="fieldData field90 displayName" style="text-align: right;">
										<span class="changeNameLink" style="text-decoration: underline; cursor: pointer; color: #0068A6; font-size: .8em;">Change</span>
									</div>
									<div class="editName" style="clear: both; float: left; display: none; text-align: right; width: 100%; padding-top: 5px;">
										<div style="float: left"><input id="mainForm:j_idt79:0:cancelEditFundingBtn" type="submit" name="mainForm:j_idt79:0:cancelEditFundingBtn" value="Cancel" class="button69grey cancelEdit" onclick="mojarra.ab(this,event,'action',0,0);return false" />
										</div>
										<div style="float: right"><input id="mainForm:j_idt79:0:saveEditFundingBtn" type="submit" name="mainForm:j_idt79:0:saveEditFundingBtn" value="Save" class="button69 saveEdit" onclick="mojarra.ab(this,event,'action','mainForm:j_idt79:0:cardName',0);return false" />
										</div>
									</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Serial Number:</div>
									<div class="fieldData field90">1202728442</div>
								</div></div>
							<div class="clear"></div>
						</div>


						<div class="whiteGreyCardBox cardData">
							<div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Type:</div>
									<div class="fieldData">ADULT</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Status:</div>
									<div class="fieldData">Active</div>
								</div>

								<div class="spacer"></div>
								<div class="infoDiv">
									<div class="fieldHeader">
										<strong>Products on Card:</strong>
									</div>
								</div>
									<div class="infoDiv">
										<div class="fieldName">Cash value:</div>
										<div class="fieldData" style="width: auto;">$89.75</div><span style="float:right; font-size:.7em;">
											<span id="d12027284420000000000000072" class="disableAutoloadSpan linkSpan">Disable Autoload</span><a id="mainForm:j_idt79:0:j_idt90:0:disableAutoloadThree" href="#" style="display:none" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:j_idt79:0:j_idt90:0:disableAutoloadThree':'mainForm:j_idt79:0:j_idt90:0:disableAutoloadThree'},'');return false">Disable Autoload</a></span>
									</div>
									<!-- Start remove tag to comment -->
									<!--  &lt;ui:remove&gt;  -->
									<div class="infoDiv">
										<div class="fieldName"><span style="float:left;">
											(Autoload amount : $ 150.00)
											</span>
										</div>
										<div class="fieldData" style="width: auto;"></div><span style="float:right; font-size:.7em;">
											<span id="dCash12027284420000000000000072" class="changeAutoloadCashAmt linkSpan">Change
												Autoload Amount</span><a id="mainForm:j_idt79:0:j_idt90:0:chnageAutoloadThree" href="#" style="display:none" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:j_idt79:0:j_idt90:0:chnageAutoloadThree':'mainForm:j_idt79:0:j_idt90:0:chnageAutoloadThree'},'');return false">Change Autoload Amount</a></span>
										<!--
										Removed Actions on Parking

										 -->

									</div>
									<!--  &lt;/ui:remove&gt;  -->
									<!-- end remove tag -->

									<div class="spacer"></div>


								<div class="infoDiv">
									<div class="fieldHeader" style="float: left;">
										<strong>Products Ready for Pick Up:</strong>
									</div>
									<div style="float: left;" class="helpToggle helpToggleInactive" onclick="showHelp('helpContainerProducts', this);"></div>

								</div><input id="mainForm:j_idt79:0:mana
52ff
ge" type="submit" name="mainForm:j_idt79:0:manage" value="Add Value" style="margin-right:2%" class="button120 bottom-right" />
								<div class="spacer10"></div>
								<div class="infoDiv">
									<div class="fieldName">View Recent Activity:</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:0:seeHistoryThirty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:0:seeHistoryThirty\':\'mainForm:j_idt79:0:seeHistoryThirty\'},\'_blank\')');return false">Last 30 Days</a>
									</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:0:seeHistorySixty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:0:seeHistorySixty\':\'mainForm:j_idt79:0:seeHistorySixty\'},\'_blank\')');return false">Last 60 Days</a>
									</div>
								</div>

							</div>
						</div>
						<div class="spacer"></div>
						<div class="darkGreyCardHeader">
							<div class="plus toggleCardInfo" style="display: block;">-</div><div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Card Nickname:</div>
									<div class="fieldData field90">
										<span class="displayName">Work</span><span class="editName" style="display: none"><input id="mainForm:j_idt79:1:cardName" type="text" name="mainForm:j_idt79:1:cardName" value="Work" /></span>
									</div>

									<div class="fieldData field90 displayName" style="text-align: right;">
										<span class="changeNameLink" style="text-decoration: underline; cursor: pointer; color: #0068A6; font-size: .8em;">Change</span>
									</div>
									<div class="editName" style="clear: both; float: left; display: none; text-align: right; width: 100%; padding-top: 5px;">
										<div style="float: left"><input id="mainForm:j_idt79:1:cancelEditFundingBtn" type="submit" name="mainForm:j_idt79:1:cancelEditFundingBtn" value="Cancel" class="button69grey cancelEdit" onclick="mojarra.ab(this,event,'action',0,0);return false" />
										</div>
										<div style="float: right"><input id="mainForm:j_idt79:1:saveEditFundingBtn" type="submit" name="mainForm:j_idt79:1:saveEditFundingBtn" value="Save" class="button69 saveEdit" onclick="mojarra.ab(this,event,'action','mainForm:j_idt79:1:cardName',0);return false" />
										</div>
									</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Serial Number:</div>
									<div class="fieldData field90">1207797539</div>
								</div></div>
							<div class="clear"></div>
						</div>


						<div class="whiteGreyCardBox cardData">
							<div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Type:</div>
									<div class="fieldData">ADULT</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Status:</div>
									<div class="fieldData">Active</div>
								</div>

								<div class="spacer"></div>
								<div class="infoDiv">
									<div class="fieldHeader">
										<strong>Products on Card:</strong>
									</div>
								</div>
									<div class="infoDiv">
										<div class="fieldName">Cash value:</div>
										<div class="fieldData" style="width: auto;">$33.90</div><span style="float:right; font-size:.7em;">
											<span id="d12077975390000000000000072" class="disableAutoloadSpan linkSpan">Disable Autoload</span><a id="mainForm:j_idt79:1:j_idt90:0:disableAutoloadThree" href="#" style="display:none" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:j_idt79:1:j_idt90:0:disableAutoloadThree':'mainForm:j_idt79:1:j_idt90:0:disableAutoloadThree'},'');return false">Disable Autoload</a></span>
									</div>
									<!-- Start remove tag to comment -->
									<!--  &lt;ui:remove&gt;  -->
									<div class="infoDiv">
										<div class="fieldName"><span style="float:left;">
											(Autoload amount : $ 50.00)
											</span>
										</div>
										<div class="fieldData" style="width: auto;"></div><span style="float:right; font-size:.7em;">
											<span id="dCash12077975390000000000000072" class="changeAutoloadCashAmt linkSpan">Change
												Autoload Amount</span><a id="mainForm:j_idt79:1:j_idt90:0:chnageAutoloadThree" href="#" style="display:none" onclick="mojarra.jsfcljs(document.getElementById('mainForm'),{'mainForm:j_idt79:1:j_idt90:0:chnageAutoloadThree':'mainForm:j_idt79:1:j_idt90:0:chnageAutoloadThree'},'');return false">Change Autoload Amount</a></span>
										<!--
										Removed Actions on Parking

										 -->

									</div>
									<!--  &lt;/ui:remove&gt;  -->
									<!-- end remove tag -->

									<div class="spacer"></div>


								<div class="infoDiv">
									<div class="fieldHeader" style="float: left;">
										<strong>Products Ready for Pick Up:</strong>
									</div>
									<div style="float: left;" class="helpToggle helpToggleInactive" onclick="showHelp('helpContainerProducts', this);"></div>

								</div><input id="mainForm:j_idt79:1:manage" type="submit" name="mainForm:j_idt79:1:manage" value="Add Value" style="margin-right:2%" class="button120 bottom-right" />
								<div class="spacer10"></div>
								<div class="infoDiv">
									<div class="fieldName">View Recent Activity:</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:1:seeHistoryThirty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:1:seeHistoryThirty\':\'mainForm:j_idt79:1:seeHistoryThirty\'},\'_blank\')');return false">Last 30 Days</a>
									</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:1:seeHistorySixty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:1:seeHistorySixty\':\'mainForm:j_idt79:1:seeHistorySixty\'},\'_blank\')');return false">Last 60 Days</a>
									</div>
								</div>

							</div>
						</div>
						<div class="spacer"></div>
						<div class="darkGreyCardHeader">
							<div class="plus toggleCardInfo" style="display: block;">-</div><div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Card Nickname:</div>
									<div class="fieldData field90">
										<span class="displayName"></span><span class="editName" style="display: none"><input id="mainForm:j_idt79:2:cardName" type="text" name="mainForm:j_idt79:2:cardName" /></span>
									</div>

									<div class="fieldData field90 displayName" style="text-align: right;">
										<span class="changeNameLink" style="text-decoration: underline; cursor: pointer; color: #0068A6; font-size: .8em;">Change</span>
									</div>
									<div class="editName" style="clear: both; float: left; display: none; text-align: right; width: 100%; padding-top: 5px;">
										<div style="float: left"><input id="mainForm:j_idt79:2:cancelEditFundingBtn" type="submit" name="mainForm:j_idt79:2:cancelEditFundingBtn" value="Cancel" class="button69grey cancelEdit" onclick="mojarra.ab(this,event,'action',0,0);return false" />
										</div>
										<div style="float: right"><input id="mainForm:j_idt79:2:saveEditFundingBtn" type="submit" name="mainForm:j_idt79:2:saveEditFundingBtn" value="Save" class="button69 saveEdit" onclick="mojarra.ab(this,event,'action','mainForm:j_idt79:2:cardName',0);return false" />
										</div>
									</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Serial Number:</div>
									<div class="fieldData field90">1201495072</div>
								</div></div>
							<div class="clear"></div>
						</div>


						<div class="whiteGreyCardBox cardData">
							<div class="cardInfo">
								<div class="infoDiv">
									<div class="fieldName">Type:</div>
									<div class="fieldData">ADULT</div>
								</div>

								<div class="infoDiv">
									<div class="fieldName">Status:</div>
									<div class="fieldData">Blocked</div>
								</div>
									<div class="infoDiv">
										<div class="fieldName">Reason:</div>
										<div class="fieldData">Lost</div>
									</div>

								<div class="spacer"></div>
								<div class="infoDiv">
									<div class="fieldHeader">
										<strong>Products on Card:</strong>
									</div>
								</div>


								<div class="infoDiv">
									<div class="fieldHeader" style="float: left;">
										<strong>Products Ready for Pick Up:</strong>
									</div>
									<div style="float: left;" class="helpToggle helpToggleInactive" onclick="showHelp('helpContainerProducts', this);"></div>

								</div>
								<div class="spacer10"></div>
								<div class="infoDiv">
									<div class="fieldName">View Recent Activity:</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:2:seeHistoryThirty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:2:seeHistoryThirty\':\'mainForm:j_idt79:2:seeHistoryThirty\'},\'_blank\')');return false">Last 30 Days</a>
									</div>
									<div class="fieldData field90"><a id="mainForm:j_idt79:2:seeHistorySixty" href="#" onclick="jsf.util.chain(this,event,'return seeHistory();','mojarra.jsfcljs(document.getElementById(\'mainForm\'),{\'mainForm:j_idt79:2:seeHistorySixty\':\'mainForm:j_idt79:2:seeHistorySixty\'},\'_blank\')');return false">Last 60 Days</a>
									</div>
								</div>

							</div>
						</div>
						<div class="spacer"></div>



					<div class="spacer"></div>
				</div>
				<!--END YOUR CLIPPER CARDS-->

				<div class="spacer"></div>
			</div>

			<!-- This is a hack to make the jquery dialog call a backend function --><span style="display:none">
				<div id="securityData"><input id="mainForm:securityQuestion" type="text" name="mainForm:securityQuestion" value="What is your city of birth?" class="securityQuestion" /><input id="mainForm:securityAnswer" type="text" name="mainForm:securityAnswer" value="Alamo, CA" class="securityAnswer" />
				</div>
				<div id="saveBtn"><input id="mainForm:saveSecurityOne" type="submit" name="mainForm:saveSecurityOne" value="Save" />
				</div></span><div id="mainForm:pnlNewEcashAmtVal" style="display:none"><input id="mainForm:newEcashAmtVal" type="text" name="mainForm:newEcashAmtVal" value="0.0" style="text-align:right;" /></div><div id="mainForm:pnlNewParkingPurseVal" style="display:none"><input id="mainForm:newParkingPurseAmtVal" type="text" name="mainForm:newParkingPurseAmtVal" value="0.0" style="text-align:right;" /></div>



			<div id="dialogGetSecurityQuestion" title="Security Question">
				<div style="font-size: .7em; text-align: left; color: #FF0000;">
					There is no security question set for this account. Select a
					security question and answer, and click “Save” to continue.<br />
				</div>
				<div class="spacer"></div>
				<div class="radioDiv" style="clear: left; font-size: .7em">
					<div class="textQuestion">
						<span class="question">Security Question</span>
					</div>
					<div class="textAnswer"><select id="mainForm:ddlSecurityQuestion" name="mainForm:ddlSecurityQuestion" class="ddlQuestion" size="1">	<option value="What is your mother's maiden name?">What is your mother's maiden name?</option>
	<option value="What is your city of birth?">What is your city of birth?</option>
	<option value="What is your best friend's name?">What is your best friend's name?</option>
</select>
					</div>
				</div>
				<div class="spacer"></div>
				<div class="radioDiv" style="clear: left; font-size: .7em">
					<div class="textQuestion">
						<span class="question">Answer</span>
					</div>
					<div class="textAnswer">
						<input maxlength="30" id="txtSecAnswer" class="input170 securityAnswer" onblur="validateAnswer(this,this.value)" />
					</div>
					<div class="valMsg"></div>
				</div>
			</div>

			<div id="dialogOptIn" title="What’s      new with Clipper?" class="dashOptDialog" style="color: #0069a6">

				<div class="spacer"></div>
				<div style="font-size: .8em; text-align: center; color: #505052">Be
					the first to know. Just edit your profile information to sign up
					for occasional email updates.</div>
				<div class="spacer"></div>
				<div>

					<a href="/ClipperCard/editProfile.jsf" style="cursor: pointer;">
						<div class="dashOpt" style="float: left; width: 28%;">Edit
							my profile now</div>
					</a><a href="#" class="closeWindow">
						<div class="dashOpt" style="width: 26%;">Remind me later</div>
					</a><a href="#" class="closeWindowAndDS">
						<div class="dashOptLeft" style="float: right; width: 26%;">No
							thanks</div>
					</a>
				</div>

			</div>


			<div id="dialogConfirmDisableAutoload" title="Confirm" style="font-size: .8em; text-align: left;">Are you sure you
				want to disable Autoload?</div>

			<div id="dialogNoHistory" title="Not Available" style="display: none">
				<span style="font-size: .75em; font-weight: normal"> You may
					not view recent activity until after 4am. </span>
			</div>

			<div id="dialogChangeAutoloadAmt" title="Change Clipper Cash Autoload Amount">
				New Amount: <input style="text-align: right" id="newEcashAmt" />
			</div>
			<div id="dialogChangeAutoloadParkingPurseAmt" title="Change Clipper Parking value Autoload Amount">
				New Amount: <input style="text-align: right" id="newParkingPurseAmt" />
			</div>
				</div>
				<div id="rightCol216">
			<!--HELP BOX-->
			<div id="helpContainerOne" class="helpContainer">
				<div class="helpArrow margin5"></div>
				<div class="helpBox">
					<div class="helpCloseX" onclick="resetHelp();"></div>
					<div class="helpTextMargin">
						<p class="helpText">
							<strong>Help: Profile Information</strong>
						</p>
						<p class="helpText">Welcome to My Clipper. You can manage your
							account information, set your card up for Autoload, add value to
							your card, and more. Use the navigation bar on the left to access
							online services.</p>
					</div>
				</div>
			</div>
			<!--END HELP BOX-->

			<!--HELP BOX-->
			<div id="helpContainerProducts" class="helpContainer margin700">
				<div class="helpArrow margin10"></div>
				<div class="helpBox">
					<div class="helpCloseX" onclick="resetHelp();"></div>
					<div class="helpTextMargin">
						<p class="helpText">
							<strong>Help: Products Ready for Pick Up</strong>
						</p>
						<p class="helpText">Certain Clipper transit products are
							loaded to your account when you tag on at the appropriate card
							reader or when the eligibility period begins.</p>
					</div>
				</div>
			</div>

			<div id="helpContainerBackup" class="helpContainer">
				<div class="helpArrow margin5"></div>
				<div class="helpBox">
					<div class="helpCloseX" onclick="resetHelp();"></div>
					<div class="helpTextMargin">

						<p class="helpText">
							<strong>Help: Backup Payment</strong>
						</p>
						<p class="helpText">
							You can add a credit card to your account profile to be used as a
							backup payment source. This backup credit card will only be
							charged if we cannot charge your primary payment source. <br />
							<br /> If you have questions about the information requested on
							this page, call Clipper Customer Service at 877-878-8883.
							(TDD/TTY 711 or 800-735-2929).

						</p>
					</div>
				</div>
				<div class="clear"></div>
			</div>
				</div>
				<div id="footer">
						<div id="footermain">
							<div id="footerleft">
								<div class="footerleftvr">
									<p class="footercustserv">
										Clipper Customer<br />Service Center
									</p>
									<p class="footercontact">
										<span class="title">Phone:</span> 877.878.8883<br /> <span class="title">TDD/TTY:</span> 711 or 800.735.2929
									</p>
									<p class="footercontact footeremail">
										<a href="mailto:custserv@clippercard.com" style="color: #000;">custserv@clippercard.com</a><br />
										<span class="title">Fax:</span> 925.686.8221
									</p>
									<p class="footercontact footerhours">
										<span class="title">Mon – Fri:</span> 6 a.m. – 8
										p.m.<br /> <span class="title">Sat – Sun:</span> 8 a.m.
										– 5 p.m.
									</p>
									<div class="footersocial" style="display: inline-block">
										<ul>
											<li class="socialIcons"><a target="_new" href="http://www.facebook.com/BayAreaClipper"><img src="/ClipperWeb/images/homepage/global/facebook_icon.png" /></a></li>
											<li class="socialIcons"><a target="_new" href="http://twitter.com/bayareaclipper"><img src="/ClipperWeb/images/homepage/global/twitter_icon.png" /></a></li>
											<li class="socialIcons"><a target="_new" href="https://www.youtube.com/channel/UCtpEI_08Af5ffdSrPEY08BQ"><img src="/ClipperWeb/images/homepage/global/youtube_icon.png" /></a></li>
										</ul>
									</div>
									<p class="footercontactforms">
										<a href="contactUs.jsf" style="color: #000;">Contact Us</a>    |    <a href="/ClipperWeb/download.do" style="color: #000;">Forms</a>
									</p>
									<hr align="left" />
									<div id="footerLogo">
										<a class="logo" href="http://www.511.org/" target="_new">Clipper</a>
									</div>

									<p></p>


								</div>
							</div>
							<div id="footercenter">
								<ul>
									<li><a href="/ClipperWeb/agreement.do" style="color: #000;">Cardholder Agreement</a></li>
									<li>|</li>
									<li><a href="https://docs.clippercard.com/brochures/en/Clipper_terms_of_use_11.14.12.pdf" target="_new" style="color: #000;">Website Terms of Use</a></li>
									<li>|</li>
									<li><a href="/ClipperWeb/privacy.do" style="color: #000;">Privacy</a></li>
									<li>|</li>
									<li><a href="/ClipperWeb/sitemap.do" style="color: #000;">Site
											Map</a></li>
								</ul>
								<hr align="left" />
								<img src="/ClipperWeb/images/homepage/global/mtc_logo.png" />
								<p class="footermtcmission">
									The Metropolitan Transportation Commission, as a public agency
									responsible for Clipper<sup style="font-size: .7em">®</sup>,
									is committed to operating its programs and services in
									accordance with federal, state and local civil rights laws and
									regulations. Please click below for more information on:
								</p>
								<p>
									<a href="http://mtc.ca.gov/about-mtc/access-everyone/ttdtty-visual-support" target="_new" style="color: #000;">Accessibility</a>
								</p>
								<p>
									<a href="http://mtc.ca.gov/about-mtc/access-everyone/civil-rights-act-file-complaint" target="_new" style="color: #000;">Non-Discrimination</a>
								</p>
								<div>
									<div style="float: left; margin-top: 10px;">
										<!-- div id=&quot;google_translate_element&quot;&gt;&lt;/div -->
									</div>
									<div>
										<ul style="padding-top: 10px;">
											<li><a href="/ClipperWeb/es/index.do" style="color: #000;">Sobre Clipper</a></li>
											<li><a href="/ClipperWeb/zh/index.do" style="color: #000;">關於 Clipper（路路通)</a></li>
										</ul>

									</div>

								</div>
								<br />
								<hr align="left" />
								<div id="footercopyright" style="margin-top: 15px;">
									<p>
										Copyright ©
										2018
										, Metropolitan Transportation Commission. All rights reserved.
									</p>
								</div>

								<div class="clear"></div>
							</div>


							<!-- END footer -->
						</div>
						<!-- END footermain -->
				</div>
			</div>

		</div><script language="javascript" type="text/javascript">//<![CDATA[
function faceletsDebug(URL) { day = new Date(); id = day.getTime(); eval("page" + id + " = window.open(URL, '" + id + "', 'toolbar=0,scrollbars=1,location=0,statusbar=0,menubar=0,resizable=1,width=800,height=600,left = 240,top = 212');"); };var faceletsOrigKeyup = document.onkeyup; document.onkeyup = function(e) { if (window.event) e = window.event; if (String.fromCharCode(e.keyCode) == 'D' & e.shiftKey & e.ctrlKey) faceletsDebug('/ClipperCard/dashboard.jsf?facelets.ui.DebugOutput=1518118886196'); else if (faceletsOrigKeyup) faceletsOrigKeyup(e); };
//]]>
</script>
		<script type="text/javascript">
			function notAlreadyProcessing() {
				!$("#loadingIcon").is(":visible")
			}
		</script><input type="hidden" name="javax.faces.ViewState" id="javax.faces.ViewState" value="3509939996826201470:-1621366080383075283" autocomplete="off" />
</form>
	<div id="loadingIcon">
		<img src="/ClipperCard/javax.faces.resource/loadSpin.gif.jsf?ln=images/loaders" />
	</div>
<script type="text/javascript">
//<![CDATA[
(function() {
var _analytics_scr = document.createElement('script');
_analytics_scr.type = 'text/javascript'; _analytics_scr.async = true; _analytics_scr.src = '/_Incapsula_Resource?SWJIYLWA=719d34d31c8e3a6e6fffd425f7e032f3&ns=3&cb=1753054381';
var _analytics_elem = document.getElementsByTagName('script')[0]; _analytics_elem.parentNode.insertBefore(_analytics_scr, _analytics_elem);
})();
// ]]>
</script></body>
</html>`)

var loginPage = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head>

	<!-- Migrated Jquery 1.4.4 to 1.8.3. Solves Security Issue: Cross-site scripting (XSS) vulnerability
	in jQuery before 1.6.3, when using location.hash to select elements, allows remote attackers to inject
	arbitrary web script or HTML via a crafted tag.	-->

	<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script>

	<base target="_parent" />

	<style type="text/css">
#myClipperModuleBody {
	margin-top: 0px;
	padding-top: 0px;
	padding-left: 5px;
}

#myclipperModuleBodyBottom {
	position: static;
	clear: left;
}

body {
	padding: 0px;
	margin: 0px;
	background: none transparent;
}

input.login {
	margin: 0px;
	border: 1px solid #C3C3C3;
}

span.text {
	color: white;
	font-size: .75em;
	padding-top: 3px;
}
</style>
	<script type="text/javascript">
		$(function() {
			$("input.login").keydown(function(e) {
				if (e.keyCode == 13) {
					e.preventDefault();
					$("input.buttonLogin").click();
				}
			});

		});
	</script><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/print2011.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/redesign2010.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/components.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/components2011.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/screen.css.jsf?ln=css" /><link type="text/css" rel="stylesheet" href="/ClipperCard/javax.faces.resource/screen2011.css.jsf?ln=css" /><script type="text/javascript" src="/ClipperCard/javax.faces.resource/jsf.js.jsf?ln=javax.faces"></script></head><body>
<form id="j_idt13" name="j_idt13" method="post" action="/ClipperCard/loginFrame.jsf" enctype="application/x-www-form-urlencoded">
<input type="hidden" name="j_idt13" value="j_idt13" />

		<div style="float: left; width: 160px;">
			<div id="myclipperModuleBody">

				<!-- start - add the ui:remove tage here to show maintenance message -->


				<div style="text-align: left; padding: 18px 0px 7px 10px;">
					<span class="text">Secure Login</span>
				</div>
				<div style="text-align: left; padding: 0px 0px 5px 10px;">
					<span class="text">Email Address:</span>
				</div>
				<div style="text-align: left; padding: 0px 0px 5px 10px;"><input id="j_idt13:username" type="text" name="j_idt13:username" class="login" />
				</div>
				<div style="text-align: left; padding: 0px 0px 5px 10px;">
					<span class="text">Password:</span>
				</div>
				<div style="text-align: left; padding: 0px 0px 7px 10px;"><input id="j_idt13:password" type="password" name="j_idt13:password" autocomplete="off" value="" maxlength="30" class="login" />
				</div>
				<span class="text"> <a href="forgottenPassword.jsf" style="color: white; font-weight: bold;" title="Forgot your password?" target="_parent">Forgot your
						password?</a>
				</span><span id="j_idt13:err"></span><input id="j_idt13:submitLogin" type="submit" name="j_idt13:submitLogin" value="Login" style="left:40px;" class="buttonLogin" onclick="mojarra.ab(this,event,'action','j_idt13:username j_idt13:password','j_idt13:err');return false" />



				<!-- end - add the ui:remove tage here to show maintenance message -->



			</div>

		</div>
		<div id="myclipperModuleBodyBottom" style="display: block; height: 20px;"></div><input type="hidden" name="javax.faces.ViewState" id="javax.faces.ViewState" value="5428792554773752026:-479288318579711101" autocomplete="off" />
</form><script type="text/javascript">
//<![CDATA[
(function() {
var _analytics_scr = document.createElement('script');
_analytics_scr.type = 'text/javascript'; _analytics_scr.async = true; _analytics_scr.src = '/_Incapsula_Resource?SWJIYLWA=719d34d31c8e3a6e6fffd425f7e032f3&ns=2&cb=374924568';
var _analytics_elem = document.getElementsByTagName('script')[0]; _analytics_elem.parentNode.insertBefore(_analytics_scr, _analytics_elem);
})();
// ]]>
</script></body>

</html>`)

var samplePages = []string{"TRANSACTION HISTORY FOR\nCARD 1202728442\nTRANSACTION TYPE\tLOCATION\tROUTE\tPRODUCT\tDEBIT\tCREDIT\tBALANCE*\n12/12/2017 09:17 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t146.85\n12/12/2017 06:21 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t144.80\n12/14/2017 09:10 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t142.75\n12/14/2017 12:08 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t140.70\n12/16/2017 04:59 PM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tBelmont\t\tClipper Cash\t12.20\t\t128.50\n12/16/2017 05:53 PM\tDual-tag exit transaction, fare adjustment (purse rebate)\t4th and King (Caltrain)\t\tClipper Cash\t\t6.75\t135.25\n12/16/2017 11:28 PM\tDual-tag entry transaction, no fare deduction\t16th St Mission\t\tClipper Cash\t\t\t135.25\n12/17/2017 12:23 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.60\t\t130.65\n12/17/2017 12:24 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t118.45\n12/17/2017 12:49 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tBelmont\t\tClipper Cash\t\t9.00\t127.45\n12/18/2017 08:02 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t115.25\n12/18/2017 08:02 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tMillbrae (Caltrain)\t\tClipper Cash\t\t12.20\t127.45\n12/18/2017 08:03 AM\tDual-tag entry transaction, no fare deduction\tMillbrae (BART)\t\tClipper Cash\t\t\t127.45\n12/18/2017 08:49 AM\tDual-tag exit transaction, fare payment\tMontgomery (BART)\t\tClipper Cash\t4.65\t\t122.80\n12/18/2017 10:22 AM\tDual-tag entry transaction, no fare deduction\tMontgomery (BART)\t\tClipper Cash\t\t\t122.80\n12/18/2017 11:13 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.65\t\t118.15\n12/18/2017 11:24 AM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tMillbrae (Caltrain)\t\tClipper Cash\t12.20\t\t105.95\n12/18/2017 11:44 AM\tDual-tag exit transaction, fare adjustment (purse rebate)\tBelmont\t\tClipper Cash\t\t9.00\t114.95\n12/20/2017 04:44 PM\tDual-tag entry transaction, maximum fare deducted (purse debit)\tBelmont\t\tClipper Cash\t12.20\t\t102.75\n12/20/2017 05:04 PM\tDual-tag exit transaction, fare adjustment (purse rebate)\tMillbrae (Caltrain)\t\tClipper Cash\t\t9.00\t111.75\n12/20/2017 05:05 PM\tDual-tag entry transaction, no fare deduction\tMillbrae (BART)\t\tClipper Cash\t\t\t111.75\n12/20/2017 05:46 PM\tDual-tag exit transaction, fare payment\tPowell St (BART)\t\tClipper Cash\t4.65\t\t107.10\n12/20/2017 07:05 PM\tSingle-tag fare payment\tPowell (Muni)\tNONE\tClipper Cash\t2.50\t\t104.60\n12/20/2017 11:33 PM\tDual-tag entry transaction, no fare deduction\tPowell St (BART)\t\tClipper Cash\t\t\t104.60\n12/21/2017 12:12 AM\tDual-tag exit transaction, fare payment\tMillbrae (BART)\t\tClipper Cash\t4.65\t\t99.95\n01/05/2018 09:02 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t97.90\n01/16/2018 08:36 PM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t95.85\n01/31/2018 08:34 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t93.80\n01/31/2018 10:37 AM\tSingle-tag fare payment\tSAM bus\tLOC\tClipper Cash\t2.05\t\t91.75\n02/10/2018\t\t\t\t\t\t\tPage 1 of", "TRANSACTION TYPE\tLOCATION\tROUTE\tPRODUCT\tDEBIT\tCREDIT\tBALANCE*\n02/01/2018 02:08 PM\tDual-tag entry transaction, no fare deduction\tMontgomery (BART)\t\tClipper Cash\t\t\t91.75\n02/01/2018 02:13 PM\tDual-tag exit transaction, fare payment\tCivic Center (BART)\t\tClipper Cash\t2.00\t\t89.75\n* If there is a discrepancy in the listing of the card balance, it may be due to a transaction not reaching the central system. Please contact the Customer Service Center at 877-878-8883 with any questions.\n02/10/2018\t\t\t\t\t\t\tPage 2 of"}
