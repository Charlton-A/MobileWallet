package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)



func TestPing(t *testing.T){
    svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   resp ,err:=http.Get(fmt.Sprintf("%s/api/v1/users/ping",ts.URL))
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   if resp.StatusCode!=401{
	   t.Fatalf("expected status code 200 got %d", resp.StatusCode)
   }
}

func TestPinAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/users/ping",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=200{
	   t.Fatalf("expected status code 200 got %d", resp.StatusCode)
   }
}

func TestUserCreateAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/users/create",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}

func TestUserBalanceAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/users/balance/1",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}


func TestUserTransferAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/users/wallet/transfer",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}

func TestUserTransactionsAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)
   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/users/transactions/1",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}

func TestWalletCreateAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/users/wallet/create",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, _:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}

func TestWalletUpdateAuth(t *testing.T){
   client := &http.Client{
		Timeout: time.Second * 10,
	}
   svr,_:=SetUp()
   ts:=httptest.NewServer(svr)

   defer ts.Close()

   req ,err:=http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/users/wallet/update",ts.URL) ,nil)
   if err!=nil{
	   t.Fatalf("expected no error got %#v",err)
   }
   req.SetBasicAuth(os.Getenv("APP_KEY"), os.Getenv("APP_PASS"))
	resp, err:= client.Do(req)
   if err != nil {
		t.Fatalf("expected no error got  %#v", err)
	}
	defer resp.Body.Close()

   if resp.StatusCode!=400{
	   t.Fatalf("expected status code 400 got %d", resp.StatusCode)
   }
}


