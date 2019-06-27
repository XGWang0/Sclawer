<html>
<head>
<title></title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
<!--===============================================================================================-->	
	<link rel="icon" type="image/png" href="/static/images/icons/favicon.ico"/>
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/static/vendor/bootstrap/css/bootstrap.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/static/fonts/font-awesome-4.7.0/css/font-awesome.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/static/vendor/animate/animate.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/static/vendor/select2/select2.min.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/vendor/perfect-scrollbar/perfect-scrollbar.css">
<!--===============================================================================================-->
	<link rel="stylesheet" type="text/css" href="/static/css/util.css">
	<link rel="stylesheet" type="text/css" href="/static/css/main.css">
<!--===============================================================================================-->
</head>
<body>

<div class="form-style-5">
<form action="/index" method="post">
<fieldset>
<legend><span class="number"></span> Searching Options</legend>
<input style="display: inline;" type="text" name="pagecount" placeholder="Page Count *">
<input style="display: inline;" type="text" name="commentcount" placeholder="Comments Count *">
<input style="display: inline;" type="text" name="votecount" placeholder="Vote *">    
<input style="display: inline;" type="submit" value="Apply" />
</fieldset>
</form>
</div>
	<div class="limiter">
		<div class="container-table100">
			<div class="wrap-table100">
				<div class="table100 ver1 m-b-110">
					<div class="table100-head">
						<table>
							<thead>
								<tr class="row100 head">
									<th class="cell100 column1">Item Name</th>
									<th class="cell100 column3">Price</th>
									<th class="cell100 column5">Positive Vote</th>
									<th class="cell100 column5">Negative Vote</th>
									<th class="cell100 column5">Vendor</th>
									<th class="cell100 column5">Comments</th>
									<th class="cell100 column3">Time</th>
								</tr>
							</thead>
						</table>
					</div>

					<!--div class="table100-body js-pscroll"-->
					<div class="table100-body">
						<table>
							<tbody>
								{{range $.itemlist}}
								<tr class="row100 body">
									<td class="cell100 column1"><a href="{{.Link}}">{{.Title}}</a></td>
									<td class="cell100 column3">{{.Price}}</td>
									<td class="cell100 column5">{{.Vote}}</td>
									<td class="cell100 column5">{{.Unvote}}</td>
									<td class="cell100 column5">{{.Vendor}}</td>
									<td class="cell100 column5"><a href="{{.CommentLink}}">{{.CommentCont}}</td>
									<td class="cell100 column3">{{.DataTime}}</td>
								</tr>
								{{end}}
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>
	</div>


<!--===============================================================================================-->	
	<script src="/static/vendor/jquery/jquery-3.2.1.min.js"></script>
<!--===============================================================================================-->
	<script src="/static/vendor/bootstrap/js/popper.js"></script>
	<script src="/static/vendor/bootstrap/js/bootstrap.min.js"></script>
<!--===============================================================================================-->
	<script src="/static/vendor/select2/select2.min.js"></script>
<!--===============================================================================================-->
	<script src="/static/vendor/perfect-scrollbar/perfect-scrollbar.min.js"></script>
	<script>
		$('.js-pscroll').each(function(){
			var ps = new PerfectScrollbar(this);

			$(window).on('resize', function(){
				ps.update();
			})
		});
			
		
	</script>
<!--===============================================================================================-->
	<script src="/static/js/main.js"></script>

</body>
</html>
