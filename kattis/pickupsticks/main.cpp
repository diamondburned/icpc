#include <bits/stdc++.h>
#include <sys/ucontext.h>
#include <unordered_map>

using namespace std;

vector<vector<int>> collect_kids(vector<vector<bool>> &loves) {
  vector<vector<int>> buses;

  auto nloves = [](vector<bool> &n) {
    int c = 0;
    for (const auto b : n) {
      if (b)
        c++;
    }
    return c;
  };

  while (true) {
    // Find the one with the most loves and bus them together.
    int max_loves = 0;
    int max_loves_idx = -1;
    for (int i = 0; i < loves.size(); i++) {
      int n = nloves(loves[i]);
      if (n > max_loves) {
        max_loves = n;
        max_loves_idx = i;
      }
    }

    // Check if we're done.
    if (max_loves_idx == -1)
      break;
  }
}

int main() {
  int nkids, npairs, buscap;
  cin >> nkids >> npairs >> buscap;

  vector<vector<bool>> hates(nkids);
  for (auto &v : hates)
    v.resize(nkids);

  unordered_map<string, int> names;
  for (int i = 0; i < nkids; i++) {
    string name;
    cin >> name;
    names[name] = i;
  }

  for (int i = 0; i < npairs; i++) {
    string a, b;
    cin >> a >> b;
    hates[names[a]][names[b]] = true;
  }

  auto loves = hates;
  for (auto &i : loves) {
    for (int ii = 0; ii < i.size(); ii++) {
      i[ii] = !i[ii];
    }
  }
}
