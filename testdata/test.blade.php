<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document Title</title>
</head>
<body>
    <p>Testing string</p>
    <p>Testing number 1245</p>
    
    <!-- Some ignoring characters test -->
    <p>Testing not found {} () string</p>
    <p>#hastag double #</p>
    <p>% Percent %</p>
    <p>_underscore _</p>
    <p>Testing dash - string</p>
    <p>Testing non-lating სატესტო</p>


    <input placeholder="Enter some text" />
    <!-- case insensitive and single-quotes -->
    <input Placeholder='Some text 2' />
    <input Placeholder='Placeholder' />
    <input placeholder="placeholder" />
    <input placeholder="Some text with 'quotes'" />

    <img alt="Alt text for image" />

    <!-- Hashtag extraction -->
    <input type="submit" value="#Send" />
    <input type="submit" value=" #value" />
    <input type="submit" value=" # Recived" />
    <p>#John with 'single'</p>
    <!-- XX it ignores this string because of duplicate check - SOLVED! -->
    <p>#John with double quotes</p>

    <!-- Title extraction -->
    <p title='John "ShotGun" Nelson'>John with double quotes</p>

    <!-- XX It removes quotes in middle of string (Shouldn't) -->
    <p title='John "ShotGun" Nelson'>John with "double" quotes</p>
    <p>#John with "double"</p>
    <input Placeholder='Some "text" - 2' />
    
</body>
</html>