package hatena

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

type mockTransport struct {
	mu        sync.Mutex
	URL       string
	Transport http.RoundTripper
}

func (t *mockTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	u, err := url.Parse(t.URL)
	if err != nil {
		return nil, err
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	r.URL = u
	if t.Transport == nil {
		t.Transport = http.DefaultTransport
	}
	return t.Transport.RoundTrip(r)
}

func TestFetchFeed(t *testing.T) {

	tests := []struct {
		name    string
		resp    []byte
		wantErr error
	}{
		{
			name: "success to get Hatena Blog Feed",
			resp: []byte(`
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="ja">
<title>emahiro/b.log</title>
<link href="https://ema-hiro.hatenablog.com/"/>
<updated>2020-03-16T11:34:12+09:00</updated>
<author>
<name>ema_hiro</name>
</author>
<generator uri="https://blog.hatena.ne.jp/" version="ea01afc657c7e490a6d01c5be9c64ba9">Hatena::Blog</generator>
<id>hatenablog://blog/12921228815712457933</id>
<entry>
	<title>エウレカに入社しました</title>
	<link href="https://ema-hiro.hatenablog.com/entry/2020/03/16/113412"/>
	<id>hatenablog://entry/26006613535851727</id>
	<published>2020-03-16T11:34:12+09:00</published>
	<updated>2020-03-16T11:34:12+09:00</updated>        <summary type="html">本日、3/16より株式会社エウレカに入社しました。 引き続き Gopher です。 ベイスターズとブレイブサンダースも変わらず応援してます。 働く場所は渋谷から赤羽橋の近くになります。 麻布十番にオフィスから歩いていけますが、渋谷と空気感が違ってまだ馴染めません。 前職より出社時刻が早くなるので「朝起きられるのか？」というのが最大の懸念事項でしたが、今のところ大丈夫そうです。 そういえば、オフィスの交差点を挟んだ向かい側に TAILORED CAFE があったので、今度利用してみようと思います。</summary>
	<content type="html">&lt;p&gt;本日、3/16より&lt;a href=&quot;https://eure.jp/&quot;&gt;株式会社エウレカ&lt;/a&gt;に入社しました。&lt;/p&gt;

&lt;p&gt;引き続き &lt;a class=&quot;keyword&quot; href=&quot;http://d.hatena.ne.jp/keyword/Gopher&quot;&gt;Gopher&lt;/a&gt; です。&lt;br /&gt;
&lt;a class=&quot;keyword&quot; href=&quot;http://d.hatena.ne.jp/keyword/%A5%D9%A5%A4%A5%B9%A5%BF%A1%BC%A5%BA&quot;&gt;ベイスターズ&lt;/a&gt;とブレイブサンダースも変わらず応援してます。&lt;/p&gt;

&lt;p&gt;働く場所は渋谷から赤羽橋の近くになります。&lt;br /&gt;
&lt;a class=&quot;keyword&quot; href=&quot;http://d.hatena.ne.jp/keyword/%CB%E3%C9%DB%BD%BD%C8%D6&quot;&gt;麻布十番&lt;/a&gt;にオフィスから歩いていけますが、渋谷と空気感が違ってまだ馴染めません。&lt;/p&gt;

&lt;p&gt;前職より出社時刻が早くなるので「朝起きられるのか？」というのが最大の懸念事項でしたが、今のところ大丈夫そうです。&lt;/p&gt;

&lt;p&gt;そういえば、オフィスの交差点を挟んだ向かい側に &lt;a href=&quot;https://tailoredcafe.jp&quot;&gt;TAILORED CAFE&lt;/a&gt; があったので、今度利用してみようと思います。&lt;/p&gt;
</content>        
	<category term="報告" label="報告" />
	
	<link rel="enclosure" href="https://cdn.blog.st-hatena.com/images/theme/og-image-1500.png" type="image/png" length="0" />

	<author>
		<name>ema_hiro</name>
	</author>
</entry>
</feed>`),
			wantErr: nil,
		},
		{
			name:    "failed to parse Hatena Blog Feed",
			resp:    []byte(`<feed `),
			wantErr: fmt.Errorf("hoge"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(tt.resp)
			}))
			defer ts.Close()

			http.DefaultTransport = &mockTransport{
				Transport: http.DefaultTransport,
				URL:       ts.URL,
			}

			if _, err := FetchFeed(); tt.wantErr == nil && err != nil {
				t.Fatalf("wantErr is %v, but actual is %v", tt.wantErr, err)
			}
		})
	}
}
