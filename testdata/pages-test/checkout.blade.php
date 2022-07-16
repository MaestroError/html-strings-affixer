<x-front-layout>

<div class="ps-checkout">
    <div class="container">
        <ul class="ps-breadcrumb">
            <li class="ps-breadcrumb__item"><a href="/">მთავარი</a></li>
            <li class="ps-breadcrumb__item active" aria-current="page"> შეძენა</li>
        </ul>
        <h3 class="ps-checkout__title"> შეძენა</h3>
        <livewire:blocks.checkout-form>
    </div>
</div>

@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>
