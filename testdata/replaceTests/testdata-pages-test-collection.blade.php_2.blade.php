<x-front-layout>
    
<div class="ps-categogy ps-categogy--separate">
    <div class="container pb-4">
        <ul class="ps-breadcrumb">
            <li class="ps-breadcrumb__item"><a href="/">{{ __('მთავარი') }}</a></li>
            <li class="ps-breadcrumb__item"><a href="{{ route('collections') }}">{{ __('ყველა კოლექცია') }}</a></li>
            <li class="ps-breadcrumb__item active" aria-current="page">{{ $collection->collection_name }}</li>
        </ul>
        <h1 class="ps-categogy__name">{{ $collection->collection_name }}</h1>

        <!-- SORTING start -->

        <!-- <div class="ps-categogy__content">
            <div class="ps-categogy__wrapper">
                <div class="ps-categogy__filter"> <a href="#" id="collapse-filter"><i class="fa fa-filter"></i><i class="fa fa-times"></i>{{ __('Filter') }}</a></div>
                <div class="ps-categogy__type"> <a href="category-list.html"><img src="/template/img/bars.png" alt></a><a href="category-grid-detail.html"><img src="/template/img/gird2.png" alt></a><a href="category-grid.html"><img src="/template/img/gird3.png" alt></a><a class="active" href="category-grid-separate.html"><img src="/template/img/gird4-active.png" alt></a></div>
                <div class="ps-categogy__sort">
                    <form><span>{{ __('Sort by') }}</span>
                        <select class="form-select">
                            <option selected>{{ __('Latest') }}</option>
                            <option value="Popularity">{{ __('Popularity') }}</option>
                            <option value="Average rating">{{ __('Average rating') }}</option>
                            <option value="Latest">{{ __('Latest') }}</option>
                            <option value="Price: low to high">{{ __('Price: low to high') }}</option>
                            <option value="Price: high to low">{{ __('Price: high to low') }}</option>
                        </select>
                    </form>
                </div>
                <div class="ps-categogy__show">
                    <form><span>{{ __('Show') }}</span>
                        <select class="form-select">
                            <option selected>12</option>
                            <option value="24">24</option>
                            <option value="36">36</option>
                            <option value="48">48</option>
                        </select>
                    </form>
                </div>
            </div>
        </div> -->

        <!-- SORTING end -->
    </div>

    
    <div class="ps-categogy__main">
        <div class="container">

            <!-- FILTER WIDGET start -->

            <!-- <div class="ps-categogy__widget"><a href="#" id="close-widget-product"><i class="fa fa-times"></i></a>
                <div class="ps-widget ps-widget--product">
                    <div class="ps-widget__block">
                        <h4 class="ps-widget__title">{{ __('Categories') }}</h4><a class="ps-block-control" href="#"><i class="fa fa-angle-down"></i></a>
                        <div class="ps-widget__content ps-widget__category">
                            <ul class="menu--mobile">
                                <li><a href="#">{{ __('Diagnosis') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Biopsy tools') }}</a></li>
                                        <li><a href="#">{{ __('Endoscopes') }}</a></li>
                                        <li><a href="#">{{ __('Monitoring') }}</a></li>
                                        <li><a href="#">{{ __('Otoscopes') }}</a></li>
                                        <li><a href="#">{{ __('Oxygen concentrators') }}</a></li>
                                        <li><a href="#">{{ __('Tables and assistants') }}</a></li>
                                        <li><a href="#">{{ __('Thermometer') }}</a></li>
                                    </ul>
                                </li>
                                <li><a href="#">{{ __('Equipment') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Blades') }}</a></li>
                                        <li><a href="#">{{ __('Education') }}</a></li>
                                        <li><a href="#">{{ __('Germicidal lamps') }}</a></li>
                                        <li><a href="#">{{ __('Heart Health') }}</a></li>
                                        <li><a href="#">{{ __('Infusion stands') }}</a></li>
                                        <li><a href="#">{{ __('Lighting') }}</a></li>
                                        <li><a href="#">{{ __('Machines') }}</a></li>
                                    </ul>
                                </li>
                                <li><a href="#">{{ __('Higiene') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Disposable products') }}</a></li>
                                        <li><a href="#">{{ __('Face masks') }}</a></li>
                                        <li><a href="#">{{ __('Gloves') }}</a></li>
                                        <li><a href="#">{{ __('Protective covers') }}</a></li>
                                        <li><a href="#">{{ __('Sterilization') }}</a></li>
                                        <li><a href="#">{{ __('Surgical foils') }}</a></li>
                                        <li><a href="#">{{ __('Uniforms') }}</a></li>
                                    </ul>
                                </li>
                                <li><a href="#">{{ __('Laboratory') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Devices') }}</a></li>
                                        <li><a href="#">{{ __('Diagnostic tests') }}</a></li>
                                        <li><a href="#">{{ __('Disinfectants') }}</a></li>
                                        <li><a href="#">{{ __('Dyes') }}</a></li>
                                        <li><a href="#">{{ __('Pipettes') }}</a></li>
                                        <li><a href="#">{{ __('Test-tubes') }}</a></li>
                                        <li><a href="#">{{ __('Tubes') }}</a></li>
                                    </ul>
                                </li>
                                <li><a href="#">{{ __('Tools') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Accessories Tools') }}</a></li>
                                        <li><a href="#">{{ __('Blood pressure') }}</a></li>
                                        <li><a href="#">{{ __('Capsules') }}</a></li>
                                        <li><a href="#">{{ __('Dental') }}</a></li>
                                        <li><a href="#">{{ __('Micrscope') }}</a></li>
                                        <li><a href="#">{{ __('Pressure') }}</a></li>
                                        <li><a href="#">{{ __('Sugar level') }}</a></li>
                                    </ul>
                                </li>
                                <li><a href="#">{{ __('Wound Care') }}</a><span class="sub-toggle"><i class="fa fa-chevron-down"></i></span>
                                    <ul class="sub-menu">
                                        <li><a href="#">{{ __('Bandages') }}</a></li>
                                        <li><a href="#">{{ __('Gypsum foundations') }}</a></li>
                                        <li><a href="#">{{ __('Patches and tapes') }}</a></li>
                                        <li><a href="#">{{ __('Stomatology') }}</a></li>
                                        <li><a href="#">{{ __('Surgical sutures') }}</a></li>
                                        <li><a href="#">{{ __('Uniforms') }}</a></li>
                                        <li><a href="#">{{ __('Wound healing') }}</a></li>
                                    </ul>
                                </li>
                            </ul>
                        </div>
                    </div>
                    <div class="ps-widget__block">
                        <h4 class="ps-widget__title">By price</h4><a class="ps-block-control" href="#"><i class="fa fa-angle-down"></i></a>
                        <div class="ps-widget__content">
                            <div class="ps-widget__price">
                                <div id="slide-price"></div>
                            </div>
                            <div class="ps-widget__input"><span class="ps-price" id="slide-price-min">$0</span><span class="bridge">-</span><span class="ps-price" id="slide-price-max">$820</span></div>
                            <button class="ps-widget__filter">{{ __('Filter') }}</button>
                        </div>
                    </div>
                    <div class="ps-widget__block">
                        <h4 class="ps-widget__title">{{ __('Color') }}</h4><a class="ps-block-control" href="#"><i class="fa fa-angle-down"></i></a>
                        <div class="ps-widget__content">
                            <div class="ps-widget__color">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorGray">
                                    <label class="custom-control-label" for="colorGray" style="--bg-color: #5b6c8f"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorGreen">
                                    <label class="custom-control-label" for="colorGreen" style="--bg-color: #12a05c"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorRed">
                                    <label class="custom-control-label" for="colorRed" style="--bg-color: #f00000"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorYellow">
                                    <label class="custom-control-label" for="colorYellow" style="--bg-color: #ff9923"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorBlack">
                                    <label class="custom-control-label" for="colorBlack" style="--bg-color: #313330"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorBlue">
                                    <label class="custom-control-label" for="colorBlue" style="--bg-color: #58c8ec"></label>
                                </div>
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="colorNavy">
                                    <label class="custom-control-label" for="colorNavy" style="--bg-color: #103178"></label>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="ps-widget__block">
                        <h4 class="ps-widget__title">{{ __('Brands') }}</h4><a class="ps-block-control" href="#"><i class="fa fa-angle-down"></i></a>
                        <div class="ps-widget__content">
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="BestPharm">
                                    <label class="custom-control-label" for="BestPharm">{{ __('BestPharm') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="HeartRate">
                                    <label class="custom-control-label" for="HeartRate">{{ __('HeartRate') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="HeartShield">
                                    <label class="custom-control-label" for="HeartShield">{{ __('HeartShield') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="iHeart">
                                    <label class="custom-control-label" for="iHeart">{{ __('iHeart') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="iLovehealth">
                                    <label class="custom-control-label" for="iLovehealth">{{ __('iLovehealth') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="Medialarm">
                                    <label class="custom-control-label" for="Medialarm">{{ __('Medialarm') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="Medicstore">
                                    <label class="custom-control-label" for="Medicstore">{{ __('Medicstore') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="MyMedi">
                                    <label class="custom-control-label" for="MyMedi">{{ __('MyMedi') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="Pharmy">
                                    <label class="custom-control-label" for="Pharmy">{{ __('Pharmy') }}</label>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="WeTakeCare">
                                    <label class="custom-control-label" for="WeTakeCare">{{ __('WeTakeCare') }}</label>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="ps-widget__block">
                        <h4 class="ps-widget__title">{{ __('Ratings') }}</h4><a class="ps-block-control" href="#"><i class="fa fa-angle-down"></i></a>
                        <div class="ps-widget__content">
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="rating5">
                                    <label class="custom-control-label" for="rating5"> </label>
                                    <div class="custom-label">
                                        <select class="ps-rating" data-read-only="true">
                                            <option value="1">1</option>
                                            <option value="2">2</option>
                                            <option value="3">3</option>
                                            <option value="4">4</option>
                                            <option value="5" selected="selected">5</option>
                                        </select><span>(6)</span>
                                    </div>
                                </div>
                            </div>
                            <div class="ps-widget__item">
                                <div class="custom-control custom-checkbox">
                                    <input class="custom-control-input" type="checkbox" id="rating4">
                                    <label class="custom-control-label" for="rating4"> </label>
                                    <div class="custom-label">
                                        <select class="ps-rating" data-read-only="true">
                                            <option value="1">1</option>
                                            <option value="2">2</option>
                                            <option value="3">3</option>
                                            <option value="4" selected="selected">4</option>
                                            <option value="5">5</option>
                                        </select><span>(9)</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="ps-widget__promo"><img src="/template/img/banner-sidebar1.jpg" alt></div>
                </div>
            </div> -->

            <!-- FILTER WIDGET end -->

            <div class="ps-categogy__product">
                <div class="row m-0">
                    @foreach($products as $product)
                        @if($product !== null)
                            <div class="col-6 col-lg-4 col-xl-3 p-0">
                                <div class="ps-product ps-product--standard">
                                    <div class="ps-product__thumbnail">
                                        <a class="ps-product__image" href="{{$product->url}}">
                                            <figure>
                                                <img src="{{$product->image}}" alt="alt" />
                                            </figure>
                                        </a>
                                        <livewire:blocks.product_actions :pid="$product->id" />
                                        <div class="ps-product__badge">
                                            @if($product->discounted >{{ __(' 0)
                                                ') }}<div class="ps-badge ps-badge--sale">{{ __('ფასდაკლება') }}</div>
                                            @endif
                                            <!-- <div class="ps-badge ps-badge--sold">{{ __('Sold Out') }}</div> -->
                                            <!-- <div class="ps-badge ps-badge--hot">Hot</div> -->
                                        </div>
                                    </div>
                                    <div class="ps-product__content">
                                        @if($product->subcategories)
                                            <a class="ps-product__branch" href="{{$product->subcategories[0]->url}}">{{$product->subcategories[0]->title}}</a>
                                        @endif
                                        <h5 class="ps-product__title">
                                            <a href="{{$product->url}}">{{$product->title}}</a>
                                        </h5>
                                        <div class="ps-product__meta">
                                            @if($product->discounted >{{ __(' 0)
                                                ') }}<span class="ps-product__price sale"> {{$product->discounted}} ₾ </span>
                                                <span class="ps-product__del">{{$product->price}} ₾</span>
                                            @else
                                                <span class="ps-product__price">{{$product->price}} ₾</span>
                                            @endif
                                        </div>
                                        <div class="ps-product__desc">
                                            <ul class="ps-product__list">
                                                @if(is_array(json_decode($product->features, true)))
                                                    @foreach (json_decode($product->features, true) as $f)
                                                        <li>{{$f['მახასიათებელი']}}</li>
                                                    @endforeach
                                                @endif
                                            </ul>
                                        </div>
                                        <livewire:blocks.product_actions :pid="$product->id" :mobile="true" :buttonCart="true">
                                    </div>
                                </div>
                            </div>
                        @endif
                    @endforeach
                    
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