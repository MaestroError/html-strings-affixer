<x-front-layout>
    
<div class="ps-compare">
    <div class="container">
        <ul class="ps-breadcrumb">
            <li class="ps-breadcrumb__item"><a href="/">{{ __('მთავარი') }}</a></li>
            <li class="ps-breadcrumb__item active" aria-current="page">{{ __('პროდუქტების შედარება') }}</li>
        </ul>
    </div>
    <div class="ps-compare__content">
        <div class="ps-compare--product">
            <div class="ps-compare__header">
                <div class="container">
                    <h2>{{ __('პროდუქტების შედარება') }}</h2>
                </div>
            </div>

            @includeIf("snippets.modals.popupCompare", ["redirect_del" => "true"])


            
        </div>
    </div>
</div>


@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>