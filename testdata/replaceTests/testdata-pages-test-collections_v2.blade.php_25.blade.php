<x-front-layout>
 


<div class="ps-categogy">
            <div class="container">
                <ul class="ps-breadcrumb">
                    <li class="ps-breadcrumb__item"><a href="index.html">{{ __('Home') }}</a></li>
                    <li class="ps-breadcrumb__item"><a href="index.html">{{ __('Shop') }}</a></li>
                    <li class="ps-breadcrumb__item active" aria-current="page">{{ __('Diagnosis') }}</li>
                </ul>
                <h1 class="ps-categogy__name">{{ __('Diagnosis') }}<sup>(32)</sup></h1>
                <div class="ps-categogy__content">
                    <div class="row row-reverse">
                        <div class="col-12 col-md-9">
                            <div class="ps-categogy__wrapper">
                                <div class="ps-categogy__type"> <a class="active" href="category-list.html"><img src="/template/img/bars-active.png" alt></a><a href="category-grid-detail.html"><img src="/template/img/gird2.png" alt></a><a href="category-grid.html"><img src="/template/img/gird3.png" alt></a><a href="category-grid-separate.html"><img src="/template/img/gird4.png" alt></a></div>
                                <div class="ps-categogy__onsale">
                                    <form>
                                        <div class="custom-control custom-checkbox">
                                            <input class="custom-control-input" type="checkbox" id="onSaleProduct">
                                            <label class="custom-control-label" for="onSaleProduct">{{ __('Show only products on sale') }}</label>
                                        </div>
                                    </form>
                                </div>
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
                            <div class="ps-categogy--list">
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/053.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('MyMedi') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Somersung Sonic X500 Basic') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(7 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price sale">$38.39</span><span class="ps-product__del">$53.99</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/054.jpg" alt="alt" /><img src="/template/img/products/057.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                                <div class="ps-badge ps-badge--hot">Hot</div>
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('MyMedi') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Somersung Sonic X2000 Pro Black') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4" selected="selected">4</option>
                                                    <option value="5">5</option>
                                                </select><span class="ps-product__review">(8 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$299.99</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/055.jpg" alt="alt" /><img src="/template/img/products/056.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('MyMedi') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Somersung Sonic X2500 Pro White') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(9 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$39.99</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/028.jpg" alt="alt" /><img src="/template/img/products/045.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('Medicstore') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Digital Thermometer X30-Pro') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4" selected="selected">4</option>
                                                    <option value="5">5</option>
                                                </select><span class="ps-product__review">(4 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price sale">$60.39</span><span class="ps-product__del">$89.99</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/042.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('iHeart') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Extractor used to remove teeth') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(5 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$53.64</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/016.jpg" alt="alt" /><img src="/template/img/products/021.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('WeTakeCare') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Oxygen concentrator model KTS-5000') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3" selected="selected">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5">5</option>
                                                </select><span class="ps-product__review">(2 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$53.99</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/001.jpg" alt="alt" /><img src="/template/img/products/009.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                                <div class="ps-badge ps-badge--sold">{{ __('Sold Out') }}</div>
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('iLovehealth') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Digital Thermometer X30-Pro') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(1 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price sale">$77.65</span><span class="ps-product__del">$80.65</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/003.jpg" alt="alt" /><img src="/template/img/products/008.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                                <div class="ps-badge ps-badge--sale">{{ __('Sale') }}</div>
                                                <div class="ps-badge ps-badge--hot">Hot</div>
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('BestPharm') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Automatic blood pressure monitor') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(3 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price sale">$90.65</span><span class="ps-product__del">$60.65</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/001.jpg" alt="alt" /><img src="/template/img/products/009.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                                <div class="ps-badge ps-badge--sold">{{ __('Sold Out') }}</div>
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('Medicstore') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Digital Thermometer X30-Pro') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(6 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price sale">$77.65</span><span class="ps-product__del">$80.65</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/011.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('BestPharm') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Hill-Rom Affinity III Progressa iBed') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4">4</option>
                                                    <option value="5" selected="selected">5</option>
                                                </select><span class="ps-product__review">(5 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$488.23</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/012.jpg" alt="alt" /><img src="/template/img/products/013.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('HeartRate') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Hill-Rom Affinity III Progressa iBed') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4" selected="selected">4</option>
                                                    <option value="5">5</option>
                                                </select><span class="ps-product__review">(3 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$436.87</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
                                <div class="ps-product ps-product--list">
                                    <div class="ps-product__content">
                                        <div class="ps-product__thumbnail"><a class="ps-product__image" href="#">
                                                <figure><img src="/template/img/products/013.jpg" alt="alt" /><img src="/template/img/products/012.jpg" alt="alt" />
                                                </figure>
                                            </a>
                                            <div class="ps-product__actions">
                                                <div class="ps-product__item" data-toggle="tooltip" data-placement="left" title="{{ __('Quick view') }}"><a href="#" data-toggle="modal" data-target="{{ __('popupQuickview') }}"><i class="fa fa-search"></i></a></div>
                                            </div>
                                            <div class="ps-product__badge">
                                            </div>
                                        </div>
                                        <div class="ps-product__info"><a class="ps-product__branch" href="#">{{ __('BestPharm') }}</a>
                                            <h5 class="ps-product__title">{{ __('<a>Hill-Rom VersaCare') }}</a></h5>
                                            <div class="ps-product__rating">
                                                <select class="ps-rating" data-read-only="true">
                                                    <option value="1">1</option>
                                                    <option value="2">2</option>
                                                    <option value="3">3</option>
                                                    <option value="4" selected="selected">4</option>
                                                    <option value="5">5</option>
                                                </select><span class="ps-product__review">(8 Reviews)</span>
                                            </div>
                                            <div class="ps-product__desc">
                                                <ul class="ps-product__list">
                                                    <li>{{ __('Study history up to 30 days') }}</li>
                                                    <li>Up to 5 users simultaneously</li>
                                                    <li>{{ __('Has HEALTH certificate') }}</li>
                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="ps-product__footer">
                                        <div class="ps-product__meta"><span class="ps-product__price">$136.87</span>
                                        </div>
                                        <div class="ps-product__quantity">
                                            <h6>{{ __('Quantity') }}</h6>
                                            <div class="def-number-input number-input safari_only">
                                                <button class="minus" onclick="this.parentNode.querySelector('input[type=number]').stepDown()"><i class="icon-minus"></i></button>
                                                <input class="quantity" min="0" name="quantity" value="1" type="number" />
                                                <button class="plus" onclick="this.parentNode.querySelector('input[type=number]').stepUp()"><i class="icon-plus"></i></button>
                                            </div><a class="ps-btn ps-btn--warning" href="#" data-toggle="modal" data-target="{{ __('popupAddcart') }}">{{ __('Add to cart') }}</a>
                                        </div>
                                        <div class="ps-product__variations"><a class="ps-product__link" href="wishlist.html">{{ __('Add to wishlist') }}</a><a class="ps-product__link" href="compare.html">{{ __('Add to Compare') }}</a></div>
                                    </div>
                                </div>
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
                        <div class="col-12 col-md-3">
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
                        </div>
                    </div>
                </div>
            </div>
</div>


@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>