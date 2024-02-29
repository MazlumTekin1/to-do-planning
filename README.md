# To Do Plannig Project

Projeyi çalıştırmak için proje ana dizininde 

```bash
  go run main.go
```
yazmanız gerekmektedir. 

Projenin arayüz kısmı yine aynı proje de static klasörünün altında bulunmaktadır. Burada index.html dosyasını açarsanız projeyi çalıştırabilirsiniz.

#### Arayüzde Developer Tasklarını Dağıt butonuna tıklanırsa hesaplama algoritması çalıştırılır ve developer'lara minimum sürede bitirilebilecek şekilde tasklar dağıtılır.

#### Developer Tasklarını Çek butonuna tıklanırsa Provider'dan veriler çekilir, eğer veriler sorunsu çekilmişse mevcut veri tabanında bulunan veriler silinir ve yeni veriler veri tabanına kaydedilir.

## Veri tabanı tabloları için DDL'ler

```bash
  CREATE TABLE test.developers (
	id int8 NOT NULL GENERATED ALWAYS AS IDENTITY( INCREMENT BY 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1 NO CYCLE),
	name varchar NULL,
	dev_work_hourly_difficult int4 NULL,
	create_date timestamp NULL,
	update_date timestamp NULL,
	is_delete bool NULL,
	is_active bool NULL,
	create_user_id varchar NULL,
	update_user_id varchar NULL
);
```

```bash
  CREATE TABLE test.tasks (
	id serial4 NOT NULL,
	"name" text NULL,
	duration int8 NULL,
	difficulty int8 NULL,
	CONSTRAINT tasks_pkey PRIMARY KEY (id)
);
```
