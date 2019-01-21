<!DOCTYPE html>

<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta name="description" content="Site Management System">
	<meta name="author" content="Aus Rasul">
	<title>Skeleton</title>
	
	<!-- outsourced css -->
	<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.4.0/css/font-awesome.min.css" rel="stylesheet" type="text/css">
	<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
	<!-- inhouse css -->
	<link rel="stylesheet" href="/public/css/base.css">
	
	<!-- outsourced scripts -->
	<script src="//code.jquery.com/jquery-latest.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.1/js/materialize.min.js"></script>
	
	<!-- inhouse js -->
	<script src="/public/js/base.js"></script>
</head>

<body>
	<div class="valign-wrapper">
		<div class="card" style="width: 25%; margin-left: auto; margin-right: auto; margin-top: 10%;">
			<div class="card-image waves-effect waves-block waves-light" style="height: 150px">
				<img class="activator" src="/public/img/google-logo-1200x630.jpg" style="max-width: 250px; max-height: 150px; margin: auto;">
			</div>
			<div class="card-content">
				<span class="card-title activator grey-text text-darken-4">Log in with Google account<i class="material-icons right">more_vert</i></span>
				<p><a href="/login/gplus">Click to log in</a></p>
			</div>
			<div class="card-reveal">
				<span class="card-title grey-text text-darken-4">Information<i class="material-icons right">close</i></span>
				<p>Logging in via your google account (a.k.a OAuth2.0) is very secure and easy.</p>
			</div>
		</div>
		<div class="card" style="width: 25%; margin-left: auto; margin-right: auto; margin-top: 10%;">
			<div class="card-image waves-effect waves-block waves-light" style="height: 150px">
				<img class="activator" src="/public/img/tim.gif" style="max-width: 250px; max-height: 150px; margin: auto;">
			</div>
			<div class="card-content">
				<span class="card-title activator grey-text text-darken-4">Log in with TIM account<i class="material-icons right">more_vert</i></span>
				<!-- <p><a href="/timlogin">Click to log in</a></p> -->
			</div>
			<div class="card-reveal">
				<span class="card-title grey-text text-darken-4">TIM Login<i class="material-icons right">close</i></span>
				<div class="row">
					<form class="col s12" method="post" action="/timlogin">
					  <div class="row">
						<div class="input-field col s6">
						  <input name="username" placeholder="Username" id="user_name" type="text" class="validate">
						  <label for="user_name">User Name</label>
						</div>
						<div class="input-field col s6">
						  <input name="password" id="password" type="password" class="validate">
						  <label for="password">Password</label>
						</div>
					  </div>
						<div class="row">
						<button class="btn waves-effect waves-light" type="submit" name="login">Login
							<i class="material-icons right">send</i>
  						</button>
						</div>
					</form>
				</div>
			</div>
		</div>
	</div>
   	
</body>
</html>
