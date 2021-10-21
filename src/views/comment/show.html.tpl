<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8">
  <meta name="viewpoint" content="width=device-width, initial-scale=1.0">
  <title>Chat App</title>
  <meta name="description" content="">
  <link rel="stylesheet" href="/static/css/style.css">
  <style type="text/css">
  </style>
  <script src="https://unpkg.com/onsenui/js/onsenui.min.js" type="text/javascript"></script>
</head>

<body>

<div class="whitebase">
    <header>
        <h1>{{.Website}}</h1>
    </header>
  
    <div class="main">

        <div class="contribute-rapper">
            <form id="post" action="/post" method="POST">
            <fieldset>
                <label for="namelabel">名前</label>
                <input type="text" name="name" id="namelabel" placeholder="例）エアグルーヴ">
            </fieldset>
            <fieldset>
                <textarea class="textarea" name="context" type="text" rows="4" cols="28" placeholder="Please post!"></textarea>
            </fieldset>
            <fieldset>
                <button type="submit" class="button">send</button>
            </fieldset>
            </form>
        </div>

        <div class="cards">
            {{range .Comment}}
            <div class="card">
                <h2 class="chat-title">{{.Id}} {{.Name}}</h2>
                <div class="chat-context">
                    {{.Context}}
                </div>
                <p class="time-stamp">{{.CreatedAt|dateformatJst}}</p>
            </div>
            {{end}}
        </div>
    </div><!-- main -->

    <footer>
        <p>[Footer]</p>   
    </footer>
</div>

<script type="text/javascript">
</script>
  
</body>
</html>