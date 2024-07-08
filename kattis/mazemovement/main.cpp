#include "BellmanFord.h"
#include <bits/extc++.h>
#include <bits/stdc++.h>

using namespace std;

int main() {
  vi nums;
  int n, r, s = -1, t = 0, si, ti;
  cin >> n;

  for (int i = 0; i < n; ++i) {
    cin >> r;
    s = min(s, r);
    if (r < s || s == -1) {
      s = r;
      si = i;
    }
    if (r > t) {
      t = r;
      ti = i;
    }
    nums.push_back(r);
  }

  for (int i = 0; i < n; ++i) {
    for (int j = i + 1; j < n; ++j) {
      int f = gcd(nums[i], nums[j]);
      if (f > 1) {
        // cout << nums[i] << ", " << nums[j] << ": " << f << endl;
        flow.addEdge(i, j, f, f);
      }
    }
  }
  auto res = flow.calc(si, ti);
  cout << res << endl;
}
