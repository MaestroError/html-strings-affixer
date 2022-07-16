<x-front-layout>
    
    
        <div class="ps-contact">
            <div class="container">
                <ul class="ps-breadcrumb">
                    <li class="ps-breadcrumb__item"><a href="/">{{ __('მთავარი') }}</a></li>
                    <li class="ps-breadcrumb__item active" aria-current="page">{{ __('კონტაქტი') }}</li>
                </ul>
                <div class="ps-contact__content">
                    <div class="row">
                        <div class="col-12 col-lg-4">
                            <div class="ps-contact__info">
                                <h2 class="ps-contact__title">{{ __('რით შეგვიძლია დახმარება?') }}</h2>
                                <p class="ps-contact__text">{{ __('ჩვენი მხარდაჭერის გუნდი ყოველთვის ხელმისაწვდომია!') }}</p>
                                <h3 class="ps-contact__fax">{{ __('- - -') }}</h3>
                                <div class="ps-contact__work">{!! App\Http\Helpers\views::getOptionValue('working_hours', $contacts) !!}</div>
                                <div class="ps-contact__email"><a href="mailto:{{ App\Http\Helpers\views::getOptionValue('email', $contacts) }}">{{ App\Http\Helpers\views::getOptionValue('email', $contacts) }}</a></div>
                                <ul class="ps-social">
                                    <li><a class="ps-social__link facebook" href="{{ App\Http\Helpers\views::getOptionValue('fb_link', $contacts) }}"><i class="fa fa-facebook"> </i><span class="ps-tooltip">{{ __('Facebook') }}</span></a></li>
                                    <li><a class="ps-social__link instagram" href="{{ App\Http\Helpers\views::getOptionValue('inst_link', $contacts) }}"><i class="fa fa-instagram"></i><span class="ps-tooltip">{{ __('Instagram') }}</span></a></li>
                                    <!-- <li><a class="ps-social__link youtube" href="#"><i class="fa fa-youtube-play"></i><span class="ps-tooltip">{{ __('Youtube') }}</span></a></li>
                                    <li><a class="ps-social__link pinterest" href="#"><i class="fa fa-pinterest-p"></i><span class="ps-tooltip">{{ __('Pinterest') }}</span></a></li>
                                    <li><a class="ps-social__link linkedin" href="#"><i class="fa fa-linkedin"></i><span class="ps-tooltip">{{ __('Linkedin') }}</span></a></li> -->
                                </ul>
                            </div>
                        </div>
                        <div class="col-12 col-lg-8">
                            <div class="ps-contact__map"><iframe src="{{ App\Http\Helpers\views::getOptionDesc('google_map_link_embed', $contacts) }}" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy"></iframe></div>
                        </div>
                    </div>
                </div>
                <form action="do_action" method="post">
                    <div class="ps-form--contact">
                        <h2 class="ps-form__title">თუ რამე კითხვები გაქვთ, შეავსეთ ფორმა და ჩვენ დაგიკავშირდებით</h2>
                        <div class="row">
                            <div class="col-12 col-md-4">
                                <div class="ps-form__group">
                                    <input class="form-control ps-form__input" type="text" placeholder="{{ __('სახელი') }}">
                                </div>
                            </div>
                            <div class="col-12 col-md-4">
                                <div class="ps-form__group">
                                    <input class="form-control ps-form__input" type="email" placeholder="{{ __('ელ-ფოსტა') }}">
                                </div>
                            </div>
                            <div class="col-12 col-md-4">
                                <div class="ps-form__group">
                                    <input class="form-control ps-form__input" type="text" placeholder="{{ __('ტელეფონი') }}">
                                </div>
                            </div>
                            <div class="col-12">
                                <div class="ps-form__group">
                                    <textarea class="form-control ps-form__textarea" rows="5" placeholder="{{ __('შეტყობინება') }}"></textarea>
                                </div>
                            </div>
                        </div>
                        <div class="ps-form__submit">
                            <button class="ps-btn ps-btn--warning">{{ __('გაგზავნა') }}</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
    

@push('modals')
 @includeif("snippets.frontModals")
@endpush

</x-front-layout>