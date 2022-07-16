<x-front-layout>

@auth
<div class="ps-wishlist">
    <div class="container">
        <ul class="ps-breadcrumb">
            <li class="ps-breadcrumb__item"><a href="/">{{ __('მთავარი') }}</a></li>
            <li class="ps-breadcrumb__item active" aria-current="page">{{ __('სურვილების სია') }}</li>
        </ul>
        <h3 class="ps-wishlist__title d-none d-md-block">{{ __('ჩემი სურვილების სია') }}</h3>
        <livewire:fav_page />
    </div>
</div>
@endauth

@guest 
    @includeIf("blocks.need_auth")
@endguest



@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>