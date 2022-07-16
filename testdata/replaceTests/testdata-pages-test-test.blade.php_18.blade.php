<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ __('Document Title') }}</title>
</head>
<body>
    <p>{{ __('Testing string') }}</p>
    <p>{{ __('Testing number 1245') }}</p>
    
    <!-- Some ignoring characters test -->
    <p>Testing not found {} () string</p>
    <p>#hastag double #</p>
    <p>% Percent %</p>
    <p>_underscore _</p>
    <p>{{ __('Testing dash - string') }}</p>
    <p>{{ __('Testing non-lating სატესტო') }}</p>


    <input placeholder="{{ __('Enter some text') }}" />
    <!-- case insensitive and single-quotes -->
    <input Placeholder='{{ __('Some text 2') }}' />
    <input Placeholder='Placeholder' />
    <input placeholder="placeholder" />

    <img alt="{{ __('Alt text for image') }}" />

    <!-- Hashtag extraction -->
    <input type="submit" value="{{ __('Send') }}" />
    <input type="submit" value=" {{ __('value') }}" />
    <input type="submit" value=" {{ __(' Recived') }}" />
    <p>{{ __('John with double') }}</p>
    <!-- XX it ignores this string because of duplicate check -->
    <p>{{ __('John with double quotes') }}</p>

    <!-- Title extraction -->
    <p title='{{ __('John "ShotGun" Nelson') }}'>{{ __('John with double quotes') }}</p>

    <!-- XX It removes quotes in middle of string (Shouldn't) -->
    <p title='{{ __('John "ShotGun" Nelson') }}'>{{ __('John with double quotes') }}</p>
    <p>{{ __('John with "double"') }}</p>
    <input Placeholder='{{ __('Some "text" - 2') }}' />
    
</body>
</html>