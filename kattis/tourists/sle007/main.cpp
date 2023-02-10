#include <algorithm>
#include <assert.h>
#include <bit>
#include <cstring>
#include <iomanip>
#include <iostream>
#include <limits.h>
#include <map>
#include <math.h>
#include <numeric>
#include <queue>
#include <set>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <utility>
#include <vector>
using namespace std;
#define ll long long
#define ld long double
#define vll vector<ll>
#define vvll vector<vll>
#define vi vector<int>
#define vvi vector<vi>
#define vvvi vector<vvi>
#define vs vector<string>
#define pi pair<int, int>
#define vpi vector<pi>
#define vvpi vector<vpi>
#define pll pair<ll, ll>
#define vpll vector<pll>
#define PB push_back
#define all(x) x.begin(), x.end()
#define inf LLONG_MAX
#define neginf LLONG_MIN
const vpi MOVES_ADJACENT{{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
// const vpi MOVES_ALL{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0,
// -1}, {1, -1}};
template <typename T> void deb(T a) { cout << a << endl; }

template <typename T> void deb(const vector<T> &v) {
  for (int i = 0; i < v.size(); ++i) {
    cout << v[i] << (i < v.size() - 1 ? " " : "");
  }
  cout << endl;
}

void deb(int *a, int n) {
  for (int i = 0; i < n; ++i) {
    cout << a[i] << " ";
  }
  cout << endl;
}

template <typename T> void deb(T *a, int n, int m) {
  for (int r = 0; r < n; ++r) {
    for (int c = 0; c < m; ++c) {
      cout << a[r][c] << " ";
    }
    cout << endl;
  }
}

vll getPrimeFacs(ll n) {
  vll p;
  while (n % 2 == 0) {
    p.push_back(2);
    n /= 2;
  }
  for (int i = 3; i * i <= n; i += 2) {
    while (n % i == 0) {
      p.push_back(i);
      n /= i;
    }
    if (n == 1) {
      break;
    }
  }
  if (n > 1) {
    p.push_back(n);
  }
  return p;
}

struct RMQ {
  ll n, k;
  vvll lookup;
  RMQ(vll &nums) {
    n = nums.size();
    k = get_k(n);
    lookup = vvll(k + 1, vll(n, LLONG_MAX));

    for (ll i = 0; i < n; ++i) {
      lookup[0][i] = nums[i];
    }

    for (ll j = 1; j < k + 1; ++j) {
      for (ll i = 0; i < n - (1 << j) + 1; ++i) {
        lookup[j][i] = min(lookup[j - 1][i], lookup[j - 1][i + (1 << (j - 1))]);
      }
    }
  }

  ll query(ll L, ll R) {
    ll ret = LLONG_MAX;
    for (ll j = k; j >= 0; --j) {
      if (L + (1 << j) - 1 <= R) {
        ret = min(ret, lookup[j][L]);
        L += 1 << j;
      }
    }
    return ret;
  }

  ll get_k(ll x) {
    ll ret = 0;
    while (x) {
      ++ret;
      x >>= 1;
    }
    return ret - 1;
  }
};

ll n;
vvll g;
vll euclid_tour;
vll heights;
map<ll, ll> height_index;

void dfs(ll u, ll parent, ll h) {
  if (height_index.count(u) == 0) {
    height_index[u] = heights.size();
  }
  heights.PB(h);
  euclid_tour.PB(u);
  for (auto v : g[u]) {
    if (v == parent) {
      continue;
    }
    dfs(v, u, h + 1);
    heights.PB(h);
    euclid_tour.PB(u);
  }
}

void solve() {
  cin >> n;
  g = vvll(n + 1, vll());
  for (ll i = 1; i < n; ++i) {
    ll u, v;
    cin >> u >> v;
    g[u].PB(v);
    g[v].PB(u);
  }

  dfs(1, -1, 1);
  RMQ rmq(heights);
  ll ret = 0;
  for (ll i = 1; i < n + 1; ++i) {
    for (ll j = i + i; j < n + 1; j += i) {
      ll a = height_index[i];
      ll b = height_index[j];
      if (a > b) {
        swap(a, b);
      }
      ll lca_height = rmq.query(a, b);
      ll a_lca = heights[a] - lca_height + 1;
      ll b_lca = heights[b] - lca_height + 1;
      if (lca_height == heights[a] || lca_height == heights[b]) {
        ret += abs(a_lca - b_lca) + 1;
      } else {
        ret += a_lca + b_lca - 1;
      }
    }
  }
  cout << ret << "\n";
}

int main() {
  ios::sync_with_stdio(0);
  // cin.tie(0);
  // ll t;
  // cin >> t;
  // while (t--)
  // {
  //     solve();
  // }
  solve();
}
