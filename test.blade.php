<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document Title</title>
</head>
<body>

    <!-- Simple string test -->
    <p>Testing string</p>
    <p>Testing number 1245</p>
    
    <!-- Some ignoring characters test -->
    <p>Testing not found {} () string</p>
    <p>#hastag double #</p>
    <p>% Percent %</p>
    <p>_underscore _</p>
    <p>Testing 'single' quotes</p>
    <p>Testing single 'quotes'</p>
    <input placeholder="Some text with 'quotes'" />
    <!-- XX Cant find last single quoted word - SOLVED! -->
    <p>#John with 'single'</p>


    <!-- Other tests -->
    <p>Testing dash - string</p>
    <p>Testing non-latin სატესტო</p>


    <!-- Placeholder test +case insensitive and single-quotes -->
    <input placeholder="Enter some text" />
    <input Placeholder='Some text 2' />

    <!-- XX replaces first placeholder  - SOLVED! (passing) -->
    <input Placeholder='test Placeholder' />
    <input placeholder="placeholder" />


    <!-- other tests -->
    <img alt="Alt text for image" />

    <!-- Hashtag extraction -->
    <input type="submit" value="#Send" />
    <input type="submit" value=" #value" />
    <input type="submit" value=" # Recived" />
    <!-- XX it ignores this string because of duplicate check - SOLVED! -->
    <p>#John with double quotes</p>

    <!-- Title extraction (with double quotes) -->
    <p title='John "ShotGun" Nelson'>John with double quotes</p>

    <!-- XX It removes quotes in middle of string (Shouldn't) - SOLVED! -->
    <p title='John "ShotGun" Nelson'>John with "double" quotes</p>

    <!-- XX Didn't removes hashtag after replace - SOLVED! -->
    <p>#John with "double"</p>

    <input Placeholder='Some "text" - 2' />

    <!-- XX Didn't finds with regex - SOLVED! (new regex) -->
    <li>Up to 5 users simultaneously</li>
    
    <!-- XX Didn't finds with regex - SOLVED! (new regex) -->
    <div class="ps-product__badge">
        <div class="ps-badge ps-badge--hot">Hot</div>
    </div>

    <!-- XX Didn't finds with regex if starting or ending with space -->
    <li> Up to 5 users simultaneously </li>


</body>
</html>