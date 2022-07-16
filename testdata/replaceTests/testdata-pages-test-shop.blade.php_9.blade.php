<x-front-layout>

<livewire:shop_main 
        :category="$category"
        :subcategory="$subcategory"
/>

@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>