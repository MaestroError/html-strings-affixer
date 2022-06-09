<x-front-layout>

<livewire:product-show 
        :product="$product"
        :category="$category"
        :subcategory="$subcategory"
/>

@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>