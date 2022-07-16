<x-front-layout>
 
<div class="ps-categogy ps-categogy--separate">
    <div class="container pb-4">
        <ul class="ps-breadcrumb">
            <li class="ps-breadcrumb__item"><a href="/">{{ __('მთავარი') }}</a></li>
            <li class="ps-breadcrumb__item active" aria-current="page">{{ __('ყველა კოლექცია') }}</li>
        </ul>
        <h1 class="ps-categogy__name">{{ __('ყველა კოლექცია') }}</h1>

       
    </div>

    
    <div class="ps-categogy__main">
        <div class="container">

        

            <div class="ps-categogy__product">
                <div class="row m-0">

                    @forelse($collections as $coll)
                    <div class="col-6 col-lg-4 col-xl-3 p-0">
                        <div class="ps-product ps-product--standard">
                            <div class="ps-product__thumbnail">
                                @isset($coll->product1->image)
                                <a class="ps-product__image" href="{{ url('collections') }}/{{ $coll->slug }}/{{ $coll->id }}">
                                    <figure>
                                        <img src="{{ $coll->product1->image }}" alt="alt" /><img src="{{ $coll->product2->image }}" alt="alt" />
                                    </figure>
                                </a>
                                @endisset
                            </div>
                            <div class="ps-product__content">
                                <h5 class="ps-product__title"><a href="{{ url('collections') }}/{{ $coll->slug }}/{{ $coll->id }}">{{ $coll->collection_name }}</a></h5>
                                @if ($coll->discounted)
                                <div class="ps-product__meta"><span class="ps-product__price sale">${{ $coll->discounted }}</span><span class="ps-product__del">${{ $coll->price }}</span>
                                @else
                                <div class="ps-product__meta"><span class="ps-product__price">${{ $coll->price }}</span>
                                @endif
                                </div>
                            </div>
                        </div>
                    </div>
                    @empty
                        <p>{{ __('კოლექციები ვერ მოიძებნა') }}</p>
                    @endforelse
                    
                </div>


                <div class="ps-pagination">
                    <ul class="pagination">
                        <li><a href="#"><i class="fa fa-angle-double-left"></i></a></li>
                        <li class="active"><a href="#">1</a></li>
                        <li><a href="#">2</a></li>
                        <li><a href="#">3</a></li>
                        <li><a href="#"><i class="fa fa-angle-double-right"></i></a></li>
                    </ul>
                </div>
                <div class="ps-delivery" data-background="/template/img/promotion/banner-delivery-2.jpg">
                    <div class="ps-delivery__content">
                        <div class="ps-delivery__text"> <i class="icon-shield-check"></i><span> <strong>100% Secure delivery </strong>{{ __('without contacting the courier') }}</span></div><a class="ps-delivery__more" href="#">{{ __('More') }}</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>