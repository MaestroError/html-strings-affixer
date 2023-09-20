<section class="refereal_" style="margin-top: 0;">
        <br>
        <div id="titletext">
          Poleć koledze -<br><br>i zbierajcie razem punkty!
        </div>
        <div id="timeline">
            <div class="dot" id="one">
                <span></span>
                <date>pierwszy krok</date>
            </div>
            <div class="dot" id="two">
                <span></span>
                <date>Drugi krok</date>
            </div>
            <div class="dot" id="three">
                <span></span>
                <date>Trzeci krok</date>
            </div>
            <div class="dot" id="four">
                <span></span>
                <date>Czwarty krok</date>
            </div>
            <div class="inside"></div>
        </div>

        <div id="titletext2">
            <center>...</center>
        </div>

        <article class='modal one'>
            <date>Użyj przycisku zaproszenia poniżej</date>
            <h2>Zaproś kolegę</h2>
            <p>
            <center><br>Wpisz <b>adres e-mail</b> lub <b>numer telefonu</b> zaproszonych współpracowników
              nauczyciel będzie mógł połączyć się z systemem. W ten sposób nowy profil zostanie natychmiast dodany
              społeczność nauczycieli KINGS, dzięki czemu nauczyciel nie będzie musiał czekać na administrację
              of KINGS, aby ją przejrzeć i zatwierdzić..<br></center>
            </p>
        </article>

        <article class='modal two'>
            <date>W razie potrzeby wyjaśnij współpracownikowi, jak się zarejestrować</date>
            <h2>Twój kolega się rejestruje</h2>
            <p>
            <center><br>Zaproszony nauczyciel z kodem polecającym może wejść na stronę KINGS (www.kingsolympiad.com)
              i wypełnij formularz rejestracyjny nauczyciela w ciągu kilku minut. Kiedy rejestrujesz nowego nauczyciela,
              musisz wprowadzić podany przez siebie kod polecenia, aby zarówno Ty, jak i Twój współpracownik mogli to zrobić
              korzystać z KRÓLÓW.<br></center>
            </p>

        </article>

        <article class='modal three'>
            <date>Uczestników Olimpiady dodaje się poprzez profil nauczyciela</date>
            <h2>Zarejestrowani uczestnicy</h2>
            <p>
            <center><br>Nauczyciel, który zarejestruje się za pomocą podanego kodu polecającego, będzie mógł się zarejestrować
              studentów na swoim profilu. Zarejestrowani studenci od razu otrzymują możliwość bezpłatnego udziału
              faza kwalifikacyjna. <br></center>
            </p>
        </article>

        <article class='modal four'>
            <date>Korzyści dla uczniów i nauczycieli!</date>
            <h2>Dodatkowe punkty!</h2>
            <p class="text-center"><br>Jeśli uczniowie zarejestrowani przez nowego nauczyciela napiszą co najmniej 10 testów do etapu kwalifikacyjnego Olimpiady, otrzymasz 40 punktów aktywności. Te dodatkowe punkty zostaną dodane do Twojego profilu na koniec okresu i będziesz mógł je wymienić na wybrane przez siebie nagrody!<br></p>
        </article>
        <div style="padding-top:0; margin-bottom: 0" class="container">
            <div class="row">
                <div class="col-12 d-flex flex-column align-items-center">
                    <div>@if($pageText) {!! $pageText->content !!} @endif</div>
                </div>
                <form id="sendViaEmailForm">
                    <input style="display: none" type="Email" placeholder="E-mail" name="email" class="form-control">
                    <button id="sendViaEmailBtn">Zaproś kolegę przez e-mail</button>
                </form>

                <form id="sendViaPhoneForm">
                    <input style="display: none" type="text" name="phone" placeholder="Telefon. nie." class="form-control">
                    <button id="sendViaPhoneBtn"> Zaproś kolegę przez telefon</button>
                </form>
            </div>
        </div>
    </section>