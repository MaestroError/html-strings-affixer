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
    <p>#hastag double #</p>
    <p>_underscore _</p>
    <!-- Some warning characters test -->
    <p>Testing not found {} () string</p>
    <p>% Percent %</p>
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
        <!-- XX Didn't finds with regex (v4) shorter then 4 -->
        <div class="ps-badge ps-badge--hot">Hot</div>
    </div>

    <!-- XX Didn't finds with regex if starting or ending with space - SOLVED! (With 3rd version of regex) -->
    <li> Up to 5 users simultaneously </li>

    <!-- No extra empty strings and mutching linebreaks with new text type suffix +TrimSpaces -->
    <li>
        Up to 5 users simultaneously
    </li>
    
    <p>Veelgestelde vragen over {{$serviceName}}</p>

    <!-- @todo Not founds/replaces first part of string  -->
    <p>Veelgestelde vragen over <span>{{$serviceName}}</span></p>

</body>
</html>